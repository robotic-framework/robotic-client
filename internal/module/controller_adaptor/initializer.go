package controller_adaptor

import (
	"github.com/eden-framework/reverse-proxy/worker"
	"github.com/robotic-framework/robotic-client/internal/constants/enums"
	"github.com/robotic-framework/robotic-client/internal/global"
	"github.com/sirupsen/logrus"
	"gobot.io/x/gobot"
)

var typeInitializer map[enums.UpstreamAdaptorType]initializer

func init() {
	typeInitializer = make(map[enums.UpstreamAdaptorType]initializer)
	typeInitializer[enums.UPSTREAM_ADAPTOR_TYPE__PROXY] = proxyAdaptorInitializer
}

type initializer func(config global.RobotConfiguration) gobot.Adaptor

func NewUpstreamAdaptor(typ enums.UpstreamAdaptorType, config global.RobotConfiguration) UpstreamAdaptor {
	initFunc, ok := typeInitializer[typ]
	if !ok {
		logrus.Panicf("cannot get initializer from upstream adaptor: %s", typ.String())
	}
	return initFunc(config)
}

type UpstreamAdaptor interface {
	gobot.Adaptor
}

func proxyAdaptorInitializer(config global.RobotConfiguration) gobot.Adaptor {
	w := &worker.Worker{
		RemoteAddr:    config.UpstreamAdaptorProxyRemoteAddr,
		RetryInterval: config.UpstreamAdaptorProxyRetryInterval,
		RetryMaxTime:  config.UpstreamAdaptorProxyRetryMaxTime,
	}
	w.Init()
	proxyAdaptor := NewProxyAdaptor(w, config.UpstreamAdaptorProxyRemotePort)
	return proxyAdaptor
}
