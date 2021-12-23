#!make
include .env
export

run:
	go run cmd/app/main.go

build:
	go build -o ./tmp/main cmd/app/main.go

format:
	go fmt ./...

lint:
	golangci-lint run

test:
	go test ./... -cover

setup:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.43.0
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

vendor:
	go mod tidy
	go mod vendor
