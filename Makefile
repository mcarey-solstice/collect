#!/usr/bin/env make

.PHONY: *

.envrc:
	cp .envrc.example .envrc

env: .envrc
	direnv allow
