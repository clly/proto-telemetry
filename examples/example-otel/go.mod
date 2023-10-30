module github.com/clly/proto-telemetry/examples/example-otel

go 1.20

replace github.com/clly/proto-telemetry => ../../

replace github.com/clly/proto-telemetry/examples/ping => ../ping

require (
	github.com/clly/proto-telemetry v0.0.0-00010101000000-000000000000
	github.com/clly/proto-telemetry/interceptor/grpc/messagemarker v0.0.0-20230620033544-dc42f9c7d107
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.42.0
	go.opentelemetry.io/otel v1.16.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.16.0
	go.opentelemetry.io/otel/sdk v1.16.0
	go.opentelemetry.io/otel/trace v1.16.0
	google.golang.org/genproto v0.0.0-20230530153820-e85fd2cbaebc
	google.golang.org/grpc v1.56.3
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	go.opentelemetry.io/otel/metric v1.16.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230526203410-71b5a4ffd15e // indirect
)
