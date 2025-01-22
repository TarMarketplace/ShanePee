SRCS   := $(shell find . -type f -name '*.go')

PREFIX := /usr/local
BINDIR := $(DESTDIR)$(PREFIX)/bin
GOBIN  := $(BINDIR)

all: api

api: $(SRCS)
	  go build -o api ./...

fmt:
		gofmt -w **/*.go

clean:
		rm -f api

