.PHONY: build run test lint install
build:
# 	bin/gendiff
	go build -o bin/gendiff ./cmd/gendiff

install:
	go install ./cmd/gendiff

run:
	go run ./cmd/gendiff/

test:
	go test -v ./...

lint:
	golangci-lint run