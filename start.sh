#!/bin/bash
if [ ! -f go.mod ]; then go mod init app; fi

go mod download
go mod verify
go get github.com/githubnemo/CompileDaemon

CompileDaemon --build="go build ." --command=./app