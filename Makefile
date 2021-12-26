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

login-db:
	psql ${DATABASE_URL}
