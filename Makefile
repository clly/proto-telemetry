SHELL = bash

BENCHFLAGS = -gcflags '-l' -benchmem -bench=. -benchtime 5s

GO_MODULE_DIRS ?= $(shell go list -m -f "{{ .Dir }}" | grep -v mod-vendor)
mkfile_path := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
GOLANGCI_CONFIG_DIR ?= $(mkfile_path)
TIMEOUT ?= 10s
GOFILES = $(shell find -type f -name '*.go' ! -name '*.pb.go')


.PHONY: dev
dev:
	./run-dev.sh

build:
	go build ./cmd/protoc-gen-go-telemetry

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint: go/lint/mod

.PHONY: tidy
tidy: go/tidy/mod

.PHONY: fmt
fmt: $(GOFILES)
	@echo "Running goimports -w"
	@goimports -w $^

.PHONY: $(GO_MODULE_DIRS)
$(GO_MODULE_DIRS):
	@echo -e "Running $(TARGET) for $(@)\n"
	make -k -f $(CURDIR)/Makefile -C $@ $(TARGET)

.PHONY: go/test/mod go/test
go/test/mod: TARGET=go/test
go/test/mod: $(GO_MODULE_DIRS)
go/test:
	@go test -timeout $(TIMEOUT) ./...

.PHONY: go/lint/mod go/lint
go/lint/mod: TARGET=go/lint fmt
go/lint/mod: $(GO_MODULE_DIRS)
go/lint:
	@golangci-lint run --config $(GOLANGCI_CONFIG_DIR)/.golangci.yml

.PHONY: go/tidy/mod go/tidy
go/tidy/mod: TARGET=go/tidy
go/tidy/mod: $(GO_MODULE_DIRS)
go/tidy:
	go mod tidy


.PHONY: bench
bench:
	go test $(BENCHFLAGS) ./test/...

.PHONY: generate
generate:
	@ ( \
		cd test ; \
		./generate.bash ; \
	);
