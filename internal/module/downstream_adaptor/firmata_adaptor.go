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

func (a *FirmataAdaptor) GetPID() (resp PIDConfig, err error) {
	return
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
