include Makefile.inc

SRCS          := $(shell find . -type f -name '*.go')

HANDLERS      := $(shell find infrastructure/handler -type f -name 'handler.go')
EMAIL         := $(shell find infrastructure/email -type f -name '*.go')
DOMAINS       := $(shell find domain -type f -name '*.go')

ORANGE        := $(shell command -v tput > /dev/null 2>&1 && tput setaf 166 || echo "")
RESET         := $(shell command -v tput > /dev/null 2>&1 && tput sgr0 || echo "")

export GOROOT := $(shell go env GOROOT)

all: out/api

out/api: cmd/api/wire_gen.go $(SRCS)
	@echo "$(ORANGE)Compiling final binary$(RESET)"
	mkdir -p out
	$(GO) build -o out ./...

cmd/api/wire_gen.go: cmd/api/wire.go cmd/api/app.go $(EMAIL)
	@echo "$(ORANGE)Hooking up dependency$(RESET)"
	$(WIRE) ./cmd/api

fmt:
	@echo "$(ORANGE)Formating$(RESET)"
	$(GOFMT) -w **/*.go

clean:
	@echo "$(ORANGE)Cleaning$(RESET)"
	rm -rf out/* cmd/api/wire_gen.go docs

docs: docs/openapi.json

docs/openapi.json: out/api
	@echo "$(ORANGE)Generating documentation$(RESET)"
	mkdir -p docs
	./out/api -command=openapi
