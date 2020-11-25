package swarm

import (
	"github.com/robotic-framework/robotic-client/internal/constants/enums"
	"github.com/robotic-framework/robotic-client/internal/global"
	"github.com/robotic-framework/robotic-client/internal/module/controller_adaptor"
	"github.com/robotic-framework/robotic-client/internal/module/downstream_adaptor"
	"github.com/sirupsen/logrus"
	"gobot.io/x/gobot"
)

type Swarm struct {
	Name string
	Mode enums.RobotMode

	// 被控驱动
	controlAdaptor controller_adaptor.UpstreamAdaptor

	// 自身下位机驱动
	downstreamAdaptor downstream_adaptor.DownstreamAdaptor
	// TODO 蜂群
	wingmen []*Wingman
	r       *gobot.Robot
}

func NewSwarm(name string, opts ...Option) *Swarm {
	r := &Swarm{
		Name: name,
	}
	r.initRobot()

	r.ApplyOption(opts...)
	return r
}

func NewSwarmFromConfig(config global.RobotConfiguration) *Swarm {
	r := &Swarm{
		Name: config.Name,
	}
	r.initRobot()

	initFunc, ok := modeInitializer[config.RobotMode]
	if !ok {
		logrus.Panicf("cannot get initializer from robot mode: %s", config.RobotMode.String())
	}
	opts := initFunc(config)
	r.ApplyOption(opts...)

	return r
}

func (r *Swarm) ApplyOption(opts ...Option) {
	for _, opt := range opts {
		opt(r)
	}
}

func (r *Swarm) initRobot() {
	r.r = gobot.NewRobot(r.Name)
}

func (r *Swarm) Start() error {
	return r.r.Start()
}

func (r *Swarm) Stop() error {
	return r.r.Stop()
}
