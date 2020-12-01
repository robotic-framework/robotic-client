package swarm

import (
	"github.com/robotic-framework/robotic-client/internal/constants/enums"
	"github.com/robotic-framework/robotic-client/internal/global"
	"github.com/robotic-framework/robotic-client/internal/module/controller_adaptor"
	"github.com/robotic-framework/robotic-client/internal/module/downstream_adaptor"
)

var modeInitializer map[enums.RobotMode]initializer

func init() {
	modeInitializer = make(map[enums.RobotMode]initializer)
	modeInitializer[enums.ROBOT_MODE__MANUAL] = manualModeInitializer
	modeInitializer[enums.ROBOT_MODE__TEST] = testModeInitializer
}

type initializer func(config global.RobotConfiguration) []Option

func manualModeInitializer(config global.RobotConfiguration) (opts []Option) {
	opts = append(opts, WithMode(enums.ROBOT_MODE__MANUAL))

	controllerAdaptor := controller_adaptor.NewUpstreamAdaptor(config.UpstreamAdaptorType, config)
	opts = append(opts, WithUpstreamAdaptor(controllerAdaptor))

	downstreamAdaptor := downstream_adaptor.NewDownstreamAdaptor(config.SelfDownstreamAdaptorType, config)
	opts = append(opts, WithSelfDownstreamAdaptor(downstreamAdaptor))
	return
}

func testModeInitializer(config global.RobotConfiguration) (opts []Option) {
	opts = append(opts, WithMode(enums.ROBOT_MODE__TEST))

	controllerAdaptor := controller_adaptor.NewUpstreamAdaptor(config.UpstreamAdaptorType, config)
	opts = append(opts, WithUpstreamAdaptor(controllerAdaptor))

	downstreamAdaptor := downstream_adaptor.NewDownstreamAdaptor(config.SelfDownstreamAdaptorType, config)
	opts = append(opts, WithSelfDownstreamAdaptor(downstreamAdaptor))

	opts = append(opts, WithWorkingFunc(TestWorkingBehavior))
	return
}
