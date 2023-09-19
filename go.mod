module github.com/ffromani/ghwadvisor

go 1.21

require (
	github.com/google/cadvisor v0.47.3
	github.com/gorilla/mux v1.8.0
	github.com/jaypipes/ghw v0.12.0
	github.com/spf13/pflag v1.0.3
	github.com/stealthrocket/net v0.2.1
	k8s.io/klog/v2 v2.100.1
)

require (
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/go-logr/logr v1.2.0 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/sys v0.3.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	howett.net/plist v1.0.0 // indirect
)

replace github.com/jaypipes/ghw => github.com/ffromani/ghw v0.12.9002
