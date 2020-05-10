.PHONY: build
build: 
	go build -v ./cmd/avitotask

.PHONY: run
run:
	go run ./cmd/avitotask/main.go

.PHONY: test
test: 
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build 