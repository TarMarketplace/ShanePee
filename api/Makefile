include Makefile.inc

SRCS     := $(shell find . -type f -name '*.go')

HANDLERS := $(shell find infrastructure/handler -type f -name '*.handler.go')

all: api

api: $(SRCS) docs/docs.go
	$(WIRE) ./cmd/api
	mkdir -p out
	$(GO) build -o out ./...

docs/docs.go: $(HANDLERS)
	$(SWAG) init -d ./cmd/api --pdl 2

docs/swagger.json: $(HANDLERS)
	$(SWAG) init -d ./cmd/api --pdl 2

docs/swagger.yaml: $(HANDLERS)
	$(SWAG) init -d ./cmd/api

fmt:
	$(GOFMT) -w **/*.go
	$(SWAG) fmt

clean:
	rm -f api

