.PHONY: build

build:
	go build ./cmd

.PHONY: run

run:
	go run ./cmd/main.go

.DEFAULT_GOAL := run