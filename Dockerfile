FROM golang:1.24-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o /goodsocd ./cmd/bot/main.go

FROM alpine:latest

RUN apk add --no-cache ca-certificates sqlite

RUN mkdir -p /data/db && mkdir -p /data/logs

COPY --from=builder /goodsocd /app/goodsocd

WORKDIR /app

VOLUME /data

CMD ["./goodsocd"]