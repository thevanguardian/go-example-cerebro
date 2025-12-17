SHELL := /bin/bash

.PHONY: tidy run test lint vuln sec tools

tidy:
	go mod tidy

tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest

run:
	go run ./cmd/api

test:
	go test ./... -race

lint:
	golangci-lint run ./...

vuln:
	govulncheck ./...

sec:
	gosec ./...
