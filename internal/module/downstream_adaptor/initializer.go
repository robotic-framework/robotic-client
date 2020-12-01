package downstream_adaptor

import (
	"github.com/robotic-framework/robotic-client/internal/constants/enums"
	"github.com/robotic-framework/robotic-client/internal/global"
	"github.com/sirupsen/logrus"
	"gobot.io/x/gobot"
)

var typeInitializer map[enums.DownstreamAdaptorType]initializer

func init() {
	typeInitializer = make(map[enums.DownstreamAdaptorType]initializer)
	typeInitializer[enums.DOWNSTREAM_ADAPTOR_TYPE__FIRMATA] = firmataAdaptorInitializer
	typeInitializer[enums.DOWNSTREAM_ADAPTOR_TYPE__MSP] = mspAdaptorInitializer
}

type initializer func(config global.RobotConfiguration) DownstreamAdaptor

func NewDownstreamAdaptor(typ enums.DownstreamAdaptorType, config global.RobotConfiguration) DownstreamAdaptor {
	initFunc, ok := typeInitializer[typ]
	if !ok {
		logrus.Panicf("cannot get initializer from downstream adaptor: %s", typ.String())
	}
	return initFunc(config)
}

type DownstreamAdaptor interface {
	gobot.Adaptor
	GetPID() (PIDConfig, error)
}

func firmataAdaptorInitializer(config global.RobotConfiguration) DownstreamAdaptor {
	firmataAdaptor := NewFirmataAdaptor(config.SelfDownstreamAdaptorFirmataName)
	return firmataAdaptor
}

func mspAdaptorInitializer(config global.RobotConfiguration) DownstreamAdaptor {
	return NewMSPAdaptor(config.SelfDownstreamAdaptorMSPName)
}
