package downstream_adaptor

import (
	"bytes"
	"gobot.io/x/gobot/platforms/firmata"
	"gobot.io/x/gobot/platforms/firmata/client"
)

type FirmataAdaptor struct {
	*firmata.Adaptor
}

func NewFirmataAdaptor(port string) *FirmataAdaptor {
	a := &FirmataAdaptor{
		Adaptor: firmata.NewAdaptor(port),
	}

	return a
}

func (a *FirmataAdaptor) GetIdentity() (resp IdentityResp, err error) {
	return
}

func (a *FirmataAdaptor) GetStatus() (resp StatusResp, err error) {
	return
}

func (a *FirmataAdaptor) GetPID() (resp PIDResp, err error) {
	return
}

func (a *FirmataAdaptor) GetAttitude() (resp AttitudeResp, err error) {
	return
}

func (a *FirmataAdaptor) GetAltitude() (resp AltitudeResp, err error) {
	return
}

func (a *FirmataAdaptor) GetRawIMU() (resp RawIMUResp, err error) {
	return
}

func (a *FirmataAdaptor) GetServo() (resp ServoResp, err error) {
	return
}

func (a *FirmataAdaptor) GetServoConfig() (resp ServoConfigResp, err error) {
	return
}

func (a *FirmataAdaptor) GetMotor() (resp MotorResp, err error) {
	return
}

func (a *FirmataAdaptor) GetMotorPins() (resp MotorPinResp, err error) {
	return
}

func (a *FirmataAdaptor) AccCalibration() error {
	return nil
}

func (a *FirmataAdaptor) MagCalibration() error {
	return nil
}

func (a *FirmataAdaptor) WriteString(str []byte) error {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteByte(client.StringData)
	for _, r := range str {
		buf.WriteByte(r & 0x7F)
		buf.WriteByte((r >> 7) & 0x7F)
	}
	return a.WriteSysex(buf.Bytes())
}
