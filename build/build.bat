@echo off
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=arm64
go build -o dp main.go

set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -o dp main.go
