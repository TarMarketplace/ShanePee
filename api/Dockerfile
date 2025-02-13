FROM golang:1.23.3 AS builder

WORKDIR /build

RUN go install github.com/swaggo/swag/cmd/swag@v1.16.4 \
    && go install github.com/google/wire/cmd/wire@v0.6.0

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN make out/api

FROM debian:bookworm-slim

COPY --from=builder /build/out/api /usr/local/bin/api

ENTRYPOINT [ "api" ]
