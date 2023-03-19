#!/usr/bin/env bash

set -eof pipefail

go install ./cmd/protoc-gen-go-telemetry/

(
    echo "Generating opentelemetry"
    cd examples/example-otel
    buf generate
)

echo 
(
    echo "Generating opencensus"
    cd examples/example-oc
    buf generate
)
