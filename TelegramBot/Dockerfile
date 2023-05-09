FROM golang:alpine AS builder

COPY . TelegramBot
WORKDIR TelegramBot

RUN go.mod download

COPY . .

RUN go build -o bot cmd\main.go

FROM alpine:latest

WORKDIR /build

COPY --from=builder TelegramBot/bot .

CMD [". /bot"]