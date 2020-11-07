FROM golang:alpine

RUN apk update && apk add --no-cache build-base bash

WORKDIR /app

COPY ./ /app

ENTRYPOINT ./start.sh