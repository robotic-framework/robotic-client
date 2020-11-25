package swarm

import "github.com/robotic-framework/robotic-client/internal/module/downstream_adaptor"

type Wingman struct {
	downstreamAdaptor *downstream_adaptor.FirmataAdaptor
}
