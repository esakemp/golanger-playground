FROM golang:alpine

RUN apk update && apk add --no-cache build-base

WORKDIR /app

COPY ./ /app

RUN go mod download
RUN go mod verify

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build ." --command=./app