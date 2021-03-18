.PHONY: build
build:
	go build -v ./cmd/TgCrypter

.DEFAULT_GOAL := build
