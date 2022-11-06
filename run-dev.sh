#!/usr/bin/env bash

set -eof pipefail

go install ./cmd/protoc-gen-go-telemetry/

(
    cd example
    buf generate
)
