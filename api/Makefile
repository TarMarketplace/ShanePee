include Makefile.inc

SRCS          := $(shell find . -type f -name '*.go')

HANDLERS      := $(shell find infrastructure/handler -type f -name '*.handler.go')
DOMAINS       := $(shell find domain -type f -name '*.go')

SWAG_FLAGS    := --parseDependency -g main.go

export GOROOT := $(shell go env GOROOT)

all: out/api

out/api: docs/docs.go cmd/api/wire_gen.go $(SRCS)
	mkdir -p out
	$(GO) build -o out ./...

cmd/api/wire_gen.go: cmd/api/wire.go cmd/api/app.go
	$(WIRE) ./cmd/api

docs/docs.go: $(HANDLERS) $(DOMAINS)
	$(SWAG) init -d ./cmd/api $(SWAG_FLAGS)

docs/swagger.json: $(HANDLERS) $(DOMAINS)
	$(SWAG) init -d ./cmd/api $(SWAG_FLAGS)

docs/swagger.yaml: $(HANDLERS) $(DOMAINS)
	$(SWAG) init -d ./cmd/api $(SWAG_FLAGS)

fmt:
	$(GOFMT) -w **/*.go
	$(SWAG) fmt

clean:
	rm -rf out/* docs/docs.go docs/swagger.json docs/swagger.yaml cmd/api/wire_gen.go

