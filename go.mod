module github.com/clly/proto-telemetry

go 1.18

require (
	github.com/clly/proto-telemetry/interceptor/grpc/messagemarker v0.0.0-00010101000000-000000000000
	github.com/davecgh/go-spew v1.1.1
	github.com/golang/protobuf v1.5.3
	github.com/google/uuid v1.3.0
	github.com/hashicorp/go-uuid v1.0.3
	github.com/stretchr/testify v1.8.2
	go.opencensus.io v0.24.0
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.41.1
	go.opentelemetry.io/otel v1.15.1
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.11.1
	go.opentelemetry.io/otel/sdk v1.14.0
	go.opentelemetry.io/otel/trace v1.15.1
	google.golang.org/genproto v0.0.0-20230306155012-7f2fa6fef1f4
	google.golang.org/grpc v1.55.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opentelemetry.io/otel/metric v0.38.1 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/clly/proto-telemetry/interceptor/grpc/messagemarker => ./interceptor/grpc/messagemarker

replace github.com/clly/proto-telemetry/examples/ping => ./examples/ping
