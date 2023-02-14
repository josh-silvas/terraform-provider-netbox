VERSION      := $(shell cat VERSION)

TEST         ?= netbox/*.go
GOFMT_FILES  ?= $$(find . -name '*.go' |grep -v vendor)

PROJECT_NAME := "terraform-provider-netbox"
COMPOSE      := docker-compose --project-name $(PROJECT_NAME) --project-directory "develop" -f "develop/docker-compose.yml"
RUN          := $(COMPOSE) run --rm develop

default: help

release:
	@$(RUN) /usr/bin/make goreleaser
.PHONY: release

go-build:
	@rm -rf ./dist/*
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./dist/${PROJECT_NAME}_${VERSION}_darwin_amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./dist/${PROJECT_NAME}_${VERSION}_linux_amd64
.PHONY: go-build

goreleaser:
	@rm -rf dist/*
	@git tag -d $(VERSION) || true
	@git tag $(VERSION)
	@VER=$(VERSION) goreleaser --rm-dist --skip-validate --skip-announce
.PHONY: goreleaser

cli:
	@$(RUN) bash
.PHONY: cli

# Run dockerized Netbox for acceptance testing
debug:
	@echo "⌛ Startup Netbox and waiting for service to become ready"
	$(COMPOSE) up --build wait
	@echo "🚀 Netbox is up and running!"
.PHONY: debug

logs:
	$(COMPOSE) logs
.PHONY: logs

destroy:
	$(COMPOSE) down --volumes
.PHONY: destroy

# -------------------------------------------------------------------------------------------
# DEVELOPMENT: Development tools for use when contributing to this project.
# -------------------------------------------------------------------------------------------
lint: .env ## Run golint on all sub-packages
	@$(RUN) make _lint
.PHONY: lint

_lint:
	@echo "Running golangci-lint..."
	@golangci-lint run --tests=false --exclude-use-default=false
.PHONY: _lint

unittest: .env ## Run UnitTest only.
	@$(RUN) make _unittest
.PHONY: unittest

_unittest: debug
	@TF_ACC=1 go test -v -short -coverprofile=coverage.txt -covermode=atomic ./... -tags=unit | { grep -v 'no test files'; true; }
.PHONY: _unittest

# -------------------------------------------------------------------------------------------
# HELPERS: Helper declarations
# -------------------------------------------------------------------------------------------
.env:
	@if [ ! -f "develop/.env" ]; then \
	   echo "Creating environment file...\nPLEASE OVERRIDE VARIABLES IN develop/.env WITH YOUR OWN VALUES!"; \
	   cp develop/example.env develop/.env; \
	fi
.PHONY: .env

help:
	@echo "\033[1m\033[01;32m\
	$(shell echo $(PROJECT_NAME) | tr  '[:lower:]' '[:upper:]') $(VERSION) \
	\033[00m\n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' \
	$(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; \
	{printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.PHONY: help