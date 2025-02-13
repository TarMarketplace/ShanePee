include Makefile.inc

SRCS          := $(shell find . -type f -name '*.go')

HANDLERS      := $(shell find infrastructure/handler -type f -name '*.handler.go')
DOMAINS       := $(shell find domain -type f -name '*.go')

export GOROOT := $(shell go env GOROOT)

all: out/api

out/api: cmd/api/wire_gen.go $(SRCS)
	mkdir -p out
	$(GO) build -o out ./...

cmd/api/wire_gen.go: cmd/api/wire.go cmd/api/app.go
	$(WIRE) ./cmd/api

fmt:
	$(GOFMT) -w **/*.go

clean:
	rm -rf out/* cmd/api/wire_gen.go

docs: all
	./out/api -command=openapi
