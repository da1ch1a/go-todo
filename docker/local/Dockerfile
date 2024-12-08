FROM golang:1.23-alpine

RUN apk update --no-cache && \
    apk add --no-cache --update git \
    mysql-client

WORKDIR /go/src/app

RUN go install github.com/cosmtrek/air@v1.50.0

COPY go.mod go.sum ./
RUN go mod download

EXPOSE 8090
