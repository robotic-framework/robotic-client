package downstream_adaptor

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"go.bug.st/serial.v1"
	"gobot.io/x/gobot"
	"io"
	"time"
)

var (
	packetHeader = []byte{HeaderStart, HeaderM, HeaderArrowReq}
)

type MSPAdaptor struct {
	name string
	port string
	conn io.ReadWriteCloser

	BoardReadyDuration time.Duration
	Dial               func(port string) (io.ReadWriteCloser, error)
	gobot.Eventer
}

func NewMSPAdaptor(port string, readyDuration time.Duration) *MSPAdaptor {
	a := &MSPAdaptor{
		name:               "MSPAdaptor",
		port:               port,
		BoardReadyDuration: readyDuration,
		Dial: func(port string) (io.ReadWriteCloser, error) {
			return serial.Open(port, &serial.Mode{
				BaudRate: 115200,
				DataBits: 8,
				Parity:   serial.NoParity,
				StopBits: serial.OneStopBit,
			})
		},
		Eventer: gobot.NewEventer(),
	}
	return a
}

func (a *MSPAdaptor) Name() string {
	return a.name
}

func (a *MSPAdaptor) SetName(n string) {
	a.name = n
}

func (a *MSPAdaptor) Connect() error {
	if a.conn == nil {
		conn, err := a.Dial(a.port)
		if err != nil {
			return err
		}
		a.conn = conn
	}

	// init board
	time.Sleep(a.BoardReadyDuration)
	logrus.Info("MSPAdaptor ready.")
	return nil
}

func (a *MSPAdaptor) Finalize() error {
	return a.conn.Close()
}

func (a *MSPAdaptor) GetIdentity() (resp IdentityResp, err error) {
	data, err := a.read(Ident)
	if err != nil {
		return
	}

	err = resp.UnmarshalBinary(data)
	return
}

func (a *MSPAdaptor) GetStatus() (resp StatusResp, err error) {
	data, err := a.read(Status)
	if err != nil {
		return
	}

	err = resp.UnmarshalBinary(data)
	return
}

func (a *MSPAdaptor) GetPID() (resp PIDResp, err error) {
	data, err := a.read(PID)
	if err != nil {
		return
	}

	err = resp.UnmarshalBinary(data)
	return
}

func (a *MSPAdaptor) GetAttitude() (resp AttitudeResp, err error) {
	data, err := a.read(Attitude)
	if err != nil {
		return
	}

	err = resp.UnmarshalBinary(data)
	return
}

func (a *MSPAdaptor) GetAltitude() (resp AltitudeResp, err error) {
	data, err := a.read(Altitude)
	if err != nil {
		return
	}

	err = resp.UnmarshalBinary(data)
	return
}

func (a *MSPAdaptor) GetRawIMU() (resp RawIMUResp, err error) {
	data, err := a.read(RawIMU)
	if err != nil {
		return
	}

	err = resp.UnmarshalBinary(data)
	return
}

func (a *MSPAdaptor) GetServo() (resp ServoResp, err error) {
	data, err := a.read(Servo)
	if err != nil {
		return
	}

	err = resp.UnmarshalBinary(data)
	return
}

func (a *MSPAdaptor) GetServoConfig() (resp ServoConfigResp, err error) {
	data, err := a.read(ServoConfig)
	if err != nil {
		return
	}

	err = resp.UnmarshalBinary(data)
	return
}

func (a *MSPAdaptor) GetMotor() (resp MotorResp, err error) {
	data, err := a.read(Motor)
	if err != nil {
		return
	}

	err = resp.UnmarshalBinary(data)
	return
}

func (a *MSPAdaptor) GetMotorPins() (resp MotorPinResp, err error) {
	data, err := a.read(MotorPins)
	if err != nil {
		return
	}

	err = resp.UnmarshalBinary(data)
	return
}

func (a *MSPAdaptor) AccCalibration() error {
	_, err := a.read(AccCalibration)
	return err
}

func (a *MSPAdaptor) MagCalibration() error {
	_, err := a.read(MagCalibration)
	return err
}

func (a *MSPAdaptor) send(dataLength, code uint8, data []byte) (int, error) {
	payload := []byte{dataLength, code}
	payload = append(payload, data...)

	buf := bytes.NewBuffer([]byte{})
	buf.Write(packetHeader)
	buf.Write(payload)

	var checksum uint8 = 0
	for _, b := range payload {
		checksum ^= b
	}
	buf.WriteByte(checksum)
	logrus.Debugf("Serial send: %x", buf.Bytes())
	return a.conn.Write(buf.Bytes())
}

func (a *MSPAdaptor) read(code uint8) ([]byte, error) {
	_, err := a.send(0, code, []byte{})
	if err != nil {
		return nil, err
	}

	err = waitHeader(a.conn)
	if err != nil {
		return nil, err
	}

	lengthAndCode := make([]byte, 2)
	_, err = a.conn.Read(lengthAndCode)
	if err != nil {
		return nil, err
	}
	length := lengthAndCode[0]

	data := make([]byte, length)
	_, err = a.conn.Read(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func waitHeader(conn io.Reader) error {
	header := make([]byte, 1)
	for {
		_, err := conn.Read(header)
		if err != nil {
			return err
		}
		logrus.Debugf("got %s", header)

		if header[0] != HeaderStart {
			continue
		}

		_, err = conn.Read(header)
		if err != nil {
			return err
		}
		logrus.Debugf("got %s", header)

		if header[0] != HeaderM {
			continue
		}

		_, err = conn.Read(header)
		if err != nil {
			return err
		}
		logrus.Debugf("got %s", header)

		if header[0] != HeaderArrowResp {
			continue
		}

		return nil
	}
}
