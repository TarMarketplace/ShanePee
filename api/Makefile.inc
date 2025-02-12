PREFIX := /usr/local
BINDIR := $(DESTDIR)$(PREFIX)/bin
GOBIN  := $(BINDIR)
GO     := go
GOFMT  := gofmt
WIRE   := wire

SWAG   := swag
# Alternatively, with docker
# SWAG   := docker run --rm -v $(pwd):/code ghcr.io/swaggo/swag:v1.8.4
# or with go
# SWAG   := go run github.com/swaggo/swag/cmd/swag
