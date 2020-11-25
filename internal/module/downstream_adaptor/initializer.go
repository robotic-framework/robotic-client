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
}

type initializer func(config global.RobotConfiguration) gobot.Adaptor

func NewDownstreamAdaptor(typ enums.DownstreamAdaptorType, config global.RobotConfiguration) DownstreamAdaptor {
	initFunc, ok := typeInitializer[typ]
	if !ok {
		logrus.Panicf("cannot get initializer from downstream adaptor: %s", typ.String())
	}
	return initFunc(config)
}

type DownstreamAdaptor interface {
	gobot.Adaptor
}

func firmataAdaptorInitializer(config global.RobotConfiguration) gobot.Adaptor {
	firmataAdaptor := NewFirmataAdaptor(config.SelfDownstreamAdaptorFirmataName)
	return firmataAdaptor
}
