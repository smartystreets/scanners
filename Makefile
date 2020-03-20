#!/usr/bin/make -f

test:
	go test -timeout=1s -race -covermode=atomic -count=1 ./...

compile:
	go build ./...

build: test compile

.PHONY: test compile build
