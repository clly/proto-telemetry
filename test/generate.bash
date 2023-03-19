#!/usr/bin/env bash

set -eof pipefail

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

