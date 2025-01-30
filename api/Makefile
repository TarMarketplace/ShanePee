include Makefile.inc

SRCS          := $(shell find . -type f -name '*.go')

HANDLERS      := $(shell find infrastructure/handler -type f -name '*.handler.go')
DOMAINS       := $(shell find domain -type f -name '*.go')

SWAG_FLAGS    := --pdl 2

export GOROOT := $(shell go env GOROOT)

all: api

api: $(SRCS) docs/docs.go
	$(WIRE) ./cmd/api
	mkdir -p out
	$(GO) build -o out ./...

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
	rm -rf out/* docs/docs.go docs/swagger.json docs/swagger.yaml

