module github.com/clly/proto-telemetry/test

go 1.20

replace github.com/clly/proto-telemetry => ../

replace github.com/clly/proto-telemetry/examples/example-otel => ../examples/example-otel

replace github.com/clly/proto-telemetry/examples/example-oc => ../examples/example-oc

require (
	github.com/clly/proto-telemetry v0.0.0-00010101000000-000000000000
	github.com/clly/proto-telemetry/examples/example-oc v0.0.0-00010101000000-000000000000
	github.com/clly/proto-telemetry/examples/example-otel v0.0.0-00010101000000-000000000000
	github.com/davecgh/go-spew v1.1.1
	github.com/google/uuid v1.3.0
	github.com/stretchr/testify v1.8.4
	go.opencensus.io v0.24.0
	go.opentelemetry.io/otel v1.16.0
	go.opentelemetry.io/otel/sdk v1.16.0
	go.opentelemetry.io/otel/trace v1.16.0
	google.golang.org/genproto v0.0.0-20230530153820-e85fd2cbaebc
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.16.0 // indirect
	go.opentelemetry.io/otel/metric v1.16.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230526203410-71b5a4ffd15e // indirect
	google.golang.org/grpc v1.56.3 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
