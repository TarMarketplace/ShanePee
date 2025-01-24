PREFIX := /usr/local
BINDIR := $(DESTDIR)$(PREFIX)/bin
GOBIN  := $(BINDIR)
GO     := go
GOFMT  := gofmt

SWAG   := swag
# Alternatively, with docker
# SWAG   := docker run --rm -v $(pwd):/code ghcr.io/swaggo/swag:latest
# or with go
# SWAG   := go run github.com/swaggo/swag/cmd/swag@latest