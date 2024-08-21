FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make build-x86

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/main /app/bin/main

COPY .env /app/.env

CMD ["/app/bin/main"]
