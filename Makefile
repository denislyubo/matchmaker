-include .env

SHELL            := /bin/sh
GOBIN            ?= $(GOPATH)/bin
PATH             := $(GOBIN):$(PATH)
GO               = go

M                = $(shell printf "\033[34;1m>>\033[0m")
TARGET_DIR       ?= $(PWD)/.build

MIGRATIONS_DIR   = ./sql/migrations/

ifeq ($(DELVE_ENABLED),true)
GCFLAGS	= -gcflags 'all=-N -l'
endif

.PHONY: start
start:
	ENV_CI=local go run ./cmd/server/main.go

.PHONY: watch
watch:
	$(info $(M) run...)
	@$(GOBIN)/air -c $(PWD)/.air.toml

.PHONY: install-tools
install-tools: $(GOBIN)
	@GOBIN=$(GOBIN) $(GO) install github.com/cosmtrek/air@latest
	@GOBIN=$(GOBIN) $(GO) install go.uber.org/mock/mockgen@latest

.PHONY: build
build:
	$(info $(M) building buysellproxy...)
	@GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(GCFLAGS) $(LDFLAGS) -o $(TARGET_DIR)/app ./cmd/server/*.go

.PHONY: db-create-migration
db-create-migration:
	$(info $(M) creating DB migration...)
	migrate create -ext sql -dir "$(MIGRATIONS_DIR)" $(filter-out $@,$(MAKECMDGOALS))

.PHONY: generate
generate:
	@$(GO) generate ./...

