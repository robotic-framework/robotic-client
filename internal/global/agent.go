package global

import config_agent "github.com/eden-framework/configuration_agent"

var AgentConfig = struct {
	Agent *config_agent.Agent
}{
	Agent: &config_agent.Agent{
		Host:               "localhost",
		Mode:               "grpc",
		Port:               8910,
		Timeout:            5,
		PullConfigInterval: 10,
		StackID:            1001,
	},
}
