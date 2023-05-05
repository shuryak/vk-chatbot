GOPATH:=$(shell go env GOPATH)

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: build
build:
	go build -o vk-chatbot cmd/app/*.go
