package main

import (
	config_agent "github.com/eden-framework/configuration_agent"
	"github.com/eden-framework/context"
	"github.com/eden-framework/eden-framework/pkg/application"
	"github.com/eden-framework/plugin-event/event"
	"github.com/robotic-framework/robotic-client/internal/global"
	"github.com/robotic-framework/robotic-client/internal/module/swarm"
)

var globalCtx *context.WaitStopContext

func main() {
	//logrus.SetLevel(logrus.DebugLevel)

	app := application.NewApplication(runner,
		application.WithConfig(&global.EventConfig, &global.AgentConfig))

	app.Start()
}

func runner(ctx *context.WaitStopContext) error {
	globalCtx = ctx

	global.AgentConfig.Agent.BindConf(&global.RobotConfig.Swarm)
	global.AgentConfig.Agent.BindBus(global.EventConfig.Event)
	global.EventConfig.Event.Subscribe(config_agent.FirstRunInitTopic, startTrigger)

	global.AgentConfig.Agent.Start()
	return nil
}

func startTrigger(m event.Message) (event.Message, error) {
	globalCtx.Add(1)
	defer globalCtx.Finish()

	r := swarm.NewSwarmFromConfig(globalCtx, global.RobotConfig.Swarm)
	go r.Start()

	<-globalCtx.Done()
	return m, r.Stop()
}
