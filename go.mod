module github.com/robotic-framework/robotic-client

go 1.14

replace (
	github.com/eden-framework/reverse-proxy => /Users/liyiwen/Documents/golang/src/github.com/eden-framework/reverse-proxy
	k8s.io/client-go => k8s.io/client-go v0.18.8
)

require (
	github.com/eden-framework/configuration_agent v0.0.0-20201124093242-e3ae61cf01ce
	github.com/eden-framework/context v0.0.3
	github.com/eden-framework/eden-framework v1.1.10-0.20201127094502-ef783952fb49
	github.com/eden-framework/enumeration v1.0.0
	github.com/eden-framework/plugin-event v0.0.5
	github.com/eden-framework/reverse-proxy v0.0.0
	github.com/gofrs/uuid v3.3.0+incompatible // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.0 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	github.com/profzone/envconfig v1.5.2
	github.com/sirupsen/logrus v1.7.0
	github.com/tarm/serial v0.0.0-20180830185346-98f6abe2eb07
	go.bug.st/serial.v1 v0.0.0-20180827123349-5f7892a7bb45
	gobot.io/x/gobot v1.14.0
	golang.org/x/crypto v0.0.0-20201117144127-c1f2f97bffc9 // indirect
)
