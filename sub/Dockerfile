# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

ADD . /go/src/app
WORKDIR /go/src/app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY sub.go ./

RUN go build -o /sub

ENTRYPOINT [ "/sub", "channels.*"]