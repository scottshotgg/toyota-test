#!/bin/bash

echo "Generating server ..."
bin/swagger_linux_amd64 generate server --name toyota-test --spec swagger/swagger.json
echo "Building server ..."
go build -mod vendor -o exe/server cmd/toyota-server/main.go 