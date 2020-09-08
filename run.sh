#!/usr/bin/env bash

echo "Pda enqueue registration documents started"
go mod vendor
rm -rf generated/*

protoc -I=./proto --go_out=generated ./proto/documents.proto

go build && go run app.go
