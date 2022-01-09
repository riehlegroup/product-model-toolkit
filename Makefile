# SPDX-FileCopyrightText: 2022 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
#
# SPDX-License-Identifier: Apache-2.0

GIT_COMMIT= `git rev-parse --short HEAD`

LDFLAGS = -ldflags "-w -s -X main.gitCommit=${GIT_COMMIT}"

.PHONY: all
all: test build

.PHONY: install
install: ## Install all Go dependencies.
	go get -v -t -d ./...

.PHONY: test
test: ## Run all tests.
	go test -race -coverprofile=coverage.out ./...

.PHONY: cover
cover: ## Show coverage from coverage.out
	go tool cover -func=coverage.out

.PHONY: build
build: build-client build-server ## Build client and server application.

.PHONY: build-client
build-client: ## Build client application.
	go build ${LDFLAGS} -o pmtclient ./cmd/client

.PHONY: build-server
build-server: ## Build server application.
	go build ${LDFLAGS} -o pmtserver ./cmd/server

.PHONY: clean
clean: ## Clean up all build artifacts.
	rm -v -f pmtclient* pmtserver* coverage.out

.PHONY: lint
lint: ## Check if code is formatted correctly.
	gofmt -d ./

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

