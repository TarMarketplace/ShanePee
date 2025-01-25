FROM golang:1.23.3 AS builder

WORKDIR /build

RUN go install github.com/swaggo/swag/cmd/swag@v1.16.4

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN make

FROM debian:bookworm-slim

COPY --from=builder /build/out/cmd /usr/local/bin/api

ENTRYPOINT [ "api" ]
