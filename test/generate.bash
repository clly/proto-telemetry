#!/usr/bin/env bash

set -eof pipefail

(
    cd ..
    go install ./cmd/protoc-gen-go-telemetry
)

(
    echo "Generating opentelemetry"
    cd open-telemetry/proto
    buf generate
)

echo 
(
    echo "Generating opencensus"
    cd opencensus/proto
    buf generate
)

