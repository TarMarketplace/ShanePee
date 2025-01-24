include Makefile.inc

SRCS     := $(shell find . -type f -name '*.go')

HANDLERS := $(shell find infrastructure/handler -type f -name '*.handler.go')

all: api

api: $(SRCS) docs/docs.go
	$(WIRE) ./cmd
	$(GO) build -o api cmd/main.go

docs/docs.go: $(HANDLERS)
	$(SWAG) init -d ./cmd --pdl 2

docs/swagger.json: $(HANDLERS)
	$(SWAG) init -d ./cmd --pdl 2

docs/swagger.yaml: $(HANDLERS)
	$(SWAG) init -d ./cmd

fmt:
	$(GOFMT) -w **/*.go
	$(SWAG) fmt

clean:
	rm -f api

