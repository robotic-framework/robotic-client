module github.com/robotic-framework/robotic-client

go 1.14

replace (
	github.com/eden-framework/reverse-proxy => /Users/liyiwen/Documents/golang/src/github.com/eden-framework/reverse-proxy
	k8s.io/client-go => k8s.io/client-go v0.18.8
)

require (
	github.com/eden-framework/configuration_agent v0.0.0-20201124093242-e3ae61cf01ce
	github.com/eden-framework/context v0.0.3
	github.com/eden-framework/eden-framework v1.1.8
	github.com/eden-framework/enumeration v1.0.0
	github.com/eden-framework/plugin-event v0.0.5
	github.com/eden-framework/reverse-proxy v0.0.0
	github.com/eden-framework/sqlx v0.0.1
	github.com/gofrs/uuid v3.3.0+incompatible // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.0 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	github.com/profzone/envconfig v1.5.1
	github.com/sirupsen/logrus v1.7.0
	gobot.io/x/gobot v1.14.0
	golang.org/x/crypto v0.0.0-20201117144127-c1f2f97bffc9 // indirect
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b // indirect
)
