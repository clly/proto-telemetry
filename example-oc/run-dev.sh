#!/usr/bin/env bash

set -euo pipefail

go build ./cmd/client
go build ./cmd/server
