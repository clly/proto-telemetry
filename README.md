# proto-telemetry

Proto-telemetry is protobuf code generator to create OpenTelemetry or 
OpenCensus helper functions to add protobuf fields to span's as attributes. 
You can more easily add observability to your existing protobuf messages and 
APIs and understand what's happening in our application without needing to 
manually trace the request and response messages. It also includes a grpc 
interceptor to automatically trace grpc requests and responses.

## Installation

```shell
go install github.com/clly/proto-telemetry@latest
```

## Example

The following example will generate a `example/example_telemetry.go` file with
the following contents:

```proto
syntax = "proto3";

package example;

import "google/protobuf/timestamp.proto";

message Example {
  string name = 1;
}
```

```go
package example

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func (x *StringMessage) TraceAttributes(ctx context.Context) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("stringmessage.name", x.Msg),
	)
}

func (x *StringMessage) TraceNamedAttributes(ctx context.Context, pfx string) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String(pfx+".stringmessage.name", x.Msg),
	)
}
```

### Extensions

proto-telemetry supports several extensions to the protobuf options to 
modify the behavior of the generated code. These extensions are defined in 
`proto/telemetry/options/v1/options.proto` and are available on the Buf 
registry at `buf.build/clly/proto-telemetry`.

#### Field Options

```protobuf
 // exclude will exclude the field from the generated code
  bool exclude = 11791;
  // field_name will override the name used in the span attribute
  string field_name = 11792;
```

#### Message Options

```protobuf
  // exclude_message will generate a function but not any span annotations. This is useful for assuming functions
// implement an interface but not generating the span annotations for the message itself.
    bool exclude_message = 11793;
// message_name will override the message prefix for fields used in the span attribute
    string message_name = 11794;
```

#### File Options

```protobuf
  // exclude_file will exclude the file from the generated code
  bool exclude_file = 11790;
```

## Usage
To use the generator, you must add the following to your `protoc` command:

```
--go-telemetry_out=paths=source_relative:. \
```

Or you can include the following in your `buf.yaml` file:

```yaml
version: v1
managed:
  enabled: true
plugins:
  - name: go
    out: gen
    opt: paths=source_relative
  - name: go-telemetry
    out: gen
    opt:
      - paths=source_relative
```
The full set of options, along with the default go generator options, can be 
seen here:
```shell

### Flags
```shell
protoc-gen-go-telemetry -h
  -include-map
    	include map key/values in trace span
  -loglevel int
    	Set the log level. Higher numbers add more logging, Tops out at 3
  -telemetry-backend string
    	Telemetry implementation to use. Supports opentelemetry or opencensus (default "opentelemetry")
```

## Benchmarks

```shell
go test -gcflags '-l' -benchmem -bench=. -benchtime 5s ./test/...
goos: linux
goarch: amd64
pkg: github.com/clly/proto-telemetry/test
cpu: Intel(R) Core(TM) i5-1035G1 CPU @ 1.00GHz
BenchmarkTraceAttributes-8   	 3354204	      1879 ns/op	     976 B/op	      16 allocs/op
BenchmarkNamedAttributes-8   	 2690252	      2236 ns/op	    1080 B/op	      20 allocs/op
PASS
ok  	github.com/clly/proto-telemetry/test	16.943s
```