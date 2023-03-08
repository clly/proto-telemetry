# proto-telemetry

This is a proto generator to create OpenTelemetry or OpenCensus helper 
functions to add span annotations for protobuf messages. 

## Usage

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