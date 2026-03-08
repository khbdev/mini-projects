FROM golang:1.25.7-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .


RUN go build -o mini-redis ./cmd/server/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/mini-redis .


CMD ["./mini-redis"]