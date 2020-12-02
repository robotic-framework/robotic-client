package downstream_adaptor

import (
	"github.com/robotic-framework/robotic-client/internal/constants/enums"
	"github.com/robotic-framework/robotic-client/internal/global"
	"github.com/sirupsen/logrus"
	"gobot.io/x/gobot"
	"time"
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
	GetIdentity() (resp IdentityResp, err error)
	GetStatus() (resp StatusResp, err error)
	GetPID() (resp PIDResp, err error)
	GetAttitude() (resp AttitudeResp, err error)
	GetAltitude() (resp AltitudeResp, err error)
	GetRawIMU() (resp RawIMUResp, err error)
	GetServo() (resp ServoResp, err error)
	GetServoConfig() (resp ServoConfigResp, err error)
	GetMotor() (resp MotorResp, err error)
	GetMotorPins() (resp MotorPinResp, err error)
	AccCalibration() error
	MagCalibration() error
}

func firmataAdaptorInitializer(config global.RobotConfiguration) DownstreamAdaptor {
	firmataAdaptor := NewFirmataAdaptor(config.SelfDownstreamAdaptorFirmataName)
	return firmataAdaptor
}

func mspAdaptorInitializer(config global.RobotConfiguration) DownstreamAdaptor {
	return NewMSPAdaptor(config.SelfDownstreamAdaptorMSPName, time.Duration(config.SelfDownstreamAdaptorMSPReadyDuration))
}
