SHELL := /bin/bash

.PHONY: run test lint build docker

run:
	API_BEARER_TOKEN?=dev-secret-token API_ADDR?:=:8080 go run ./cmd/api

test:
	go test ./... -race -count=1

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/api ./cmd/api

docker:
	docker build -t scheduling-api:dev .
