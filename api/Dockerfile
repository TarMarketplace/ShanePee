FROM golang:1.23.3 AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN make

FROM debian:bookworm-slim

COPY --from=builder /build/api /usr/local/bin/api

ENTRYPOINT [ "api" ]
