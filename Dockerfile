FROM golang:1.21 AS builder
MAINTAINER Artem Smirnov <liveartem@yandex.ru>

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/app

FROM alpine:latest

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

CMD ["./main"]