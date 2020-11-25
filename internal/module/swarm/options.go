package swarm

import (
	"github.com/robotic-framework/robotic-client/internal/constants/enums"
	"github.com/robotic-framework/robotic-client/internal/module/controller_adaptor"
	"github.com/robotic-framework/robotic-client/internal/module/downstream_adaptor"
)

type Option func(r *Swarm)

func WithUpstreamAdaptor(adaptor controller_adaptor.UpstreamAdaptor) Option {
	return func(r *Swarm) {
		r.controlAdaptor = adaptor
		r.r.AddConnection(adaptor)
	}
}

func WithSelfDownstreamAdaptor(adaptor downstream_adaptor.DownstreamAdaptor) Option {
	return func(r *Swarm) {
		r.downstreamAdaptor = adaptor
		r.r.AddConnection(adaptor)
	}
}

func WithManualMode(adaptor *controller_adaptor.ProxyAdaptor) Option {
	return func(r *Swarm) {
		r.Mode = enums.ROBOT_MODE__MANUAL
		r.controlAdaptor = adaptor
		r.r.AddConnection(adaptor)
	}
}

func WithSelfTestMode() Option {
	return func(r *Swarm) {
		r.Mode = enums.ROBOT_MODE__TEST
		r.r.Work = SelfTestModeWork
	}
}

func SelfTestModeWork() {

}
