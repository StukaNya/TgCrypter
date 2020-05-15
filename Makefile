.PHONY: build
build:
	go build -v ./cmd/steamrest

.DEFAULT_GOAL := build
