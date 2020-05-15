.PHONY: install
install: 
	go install ./cmd/avitotask

.PHONY: build
build: 
	go build ./cmd/avitotask

.PHONY: run
run:
	go run ./cmd/avitotask/main.go

.PHONY: test
test: 
	go test -v -race ./...

.DEFAULT_GOAL := build 