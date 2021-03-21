.PHONY: build
build:
	go build -v ./cmd/tg-crypter

.DEFAULT_GOAL := build
