include Makefile.inc

SRCS          := $(shell find . -type f -name '*.go')

HANDLERS      := $(shell find infrastructure/handler -type f -name '*.handler.go')
EMAIL         := $(shell find infrastructure/email -type f -name '*.go')
DOMAINS       := $(shell find domain -type f -name '*.go')

export GOROOT := $(shell go env GOROOT)

all: out/api

out/api: cmd/api/wire_gen.go $(SRCS)
	mkdir -p out
	$(GO) build -o out ./...

cmd/api/wire_gen.go: cmd/api/wire.go cmd/api/app.go $(EMAIL)
	$(WIRE) ./cmd/api

fmt:
	$(GOFMT) -w **/*.go

clean:
	rm -rf out/* cmd/api/wire_gen.go docs

docs: out/api
	mkdir -p docs
	./out/api -command=openapi
