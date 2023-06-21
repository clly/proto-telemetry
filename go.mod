module github.com/clly/proto-telemetry

go 1.18

require (
	github.com/golang/protobuf v1.5.3
	github.com/hashicorp/go-uuid v1.0.3
	github.com/stretchr/testify v1.8.3
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/clly/proto-telemetry/interceptor/grpc/messagemarker => ./interceptor/grpc/messagemarker

replace github.com/clly/proto-telemetry/examples/example-oc => ./examples/example-oc

replace github.com/clly/proto-telemetry/examples/example-otel => ./examples/example-otel

replace github.com/clly/proto-telemetry/examples/ping => ./examples/ping
