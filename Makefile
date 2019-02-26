#!/usr/bin/env make

.PHONY: *

.envrc:
	cp .envrc.example .envrc

env: .envrc
	direnv allow

bin:
	mkdir -p bin

bin/collect:
	go build -o bin/collect
