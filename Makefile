SHELL = bash

BENCHFLAGS = -gcflags '-l' -benchmem -bench=. -benchtime 5s

.PHONY: dev
dev:
	./run-dev.sh

.PHONY: tests
tests:
	go test ./...

.PHONY: bench
bench:
	go test $(BENCHFLAGS) ./test/...

.PHONY: generate
generate:
	@ ( \
		cd test ; \
		./generate.bash ; \
	);