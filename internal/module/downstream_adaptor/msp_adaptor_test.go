package downstream_adaptor

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestGetPID(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	adaptor := NewMSPAdaptor("/dev/cu.usbmodem143101")
	err := adaptor.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer adaptor.Finalize()

	for {
		pid, err := adaptor.GetPID()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("%+v\n", pid)
	}
}
