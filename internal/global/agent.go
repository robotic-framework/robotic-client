package global

import config_agent "github.com/eden-framework/configuration_agent"

var AgentConfig = struct {
	Agent *config_agent.Agent
}{
	Agent: &config_agent.Agent{
		Host:               "127.0.0.1",
		Mode:               "grpc",
		Port:               8900,
		Timeout:            5,
		PullConfigInterval: 10,
		StackID:            0,
	},
}
