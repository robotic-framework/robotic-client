package downstream_adaptor

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
	"gobot.io/x/gobot"
	"io"
)

var (
	packetHeader = []byte{HeaderStart, HeaderM, HeaderArrowReq}
)

type MSPAdaptor struct {
	name string
	port string
	conn io.ReadWriteCloser
	Dial func(port string) (io.ReadWriteCloser, error)
	gobot.Eventer
}

func NewMSPAdaptor(port string) *MSPAdaptor {
	a := &MSPAdaptor{
		name: "MSPAdaptor",
		port: port,
		Dial: func(port string) (io.ReadWriteCloser, error) {
			return serial.OpenPort(&serial.Config{
				Name:     port,
				Baud:     115200,
				Size:     8,
				Parity:   serial.ParityNone,
				StopBits: serial.Stop1,
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
	return nil
}

func (a *MSPAdaptor) Finalize() error {
	return a.conn.Close()
}

func (a *MSPAdaptor) GetPID() (resp PIDConfig, err error) {
	data, err := a.read(PID)
	if err != nil {
		return
	}

	err = resp.UnmarshalBinary(data)
	return
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

	// waiting correct header
	header := make([]byte, 3)
	for {
		_, err = a.conn.Read(header)
		if err != nil {
			return nil, err
		}

		if header[0] == HeaderStart && header[1] == HeaderM && header[2] == HeaderArrowResp {
			break
		}
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
