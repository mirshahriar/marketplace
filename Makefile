PROJECT_NAME := "github.com/mirshahriar/marketplace"
PKG := "github.com/mirshahriar/$(PROJECT_NAME)"

GO ?= $(shell command -v go 2> /dev/null)
PACKAGES=$(shell go list ./...)

TOOLS_BIN_DIR := $(abspath bin)
GO_INSTALL = ./scripts/go_install.sh

GOIMPORTS_VER := master
GOIMPORTS_BIN := goimports
GOIMPORTS := $(TOOLS_BIN_DIR)/$(GOIMPORTS_BIN)

GODOC_VER := master
GODOC_BIN := gomarkdoc
GODOC := $(TOOLS_BIN_DIR)/$(GODOC_BIN)

GOLANGCILINT_VER := v1.53.1
GOLANGCILINT_BIN := golangci-lint
GOLANGCILINT := $(TOOLS_BIN_DIR)/$(GOLANGCILINT_BIN)

.PHONY: all dep test build check goimports goformat prepare docker

all: build

migrate:
	@go run main.go migrate

run:
	@go run main.go serve

test: ## Run unittests
	@rm -f ${DB_LIST}
	@go clean -testcache
	@go test -p 20 -timeout 1800s -failfast -short ${PACKAGES}

dep: ## Get the dependencies
	@go mod vendor
	@go mod tidy

build: ## Build the binary file
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GO) build -a -installsuffix cgo -o binary/product-management  .

clean: ## Remove previous build
	@rm -f ${DB_LIST}
	@rm -f binary/$(PROJECT_NAME)

lint: $(GOLANGCILINT)
	@echo Running golangci-lint
	$(GOLANGCILINT) run

govet:
	@echo Running govet
	$(GO) vet ./...
	@echo Govet success

goimports: $(GOIMPORTS)
	@echo Checking if imports are sorted
	@for package in $(PACKAGES); do \
		echo "Checking "$$package; \
		files=$$(go list -f '{{range .GoFiles}}{{$$.Dir}}/{{.}} {{end}}' $$package); \
		if [ "$$files" ]; then \
			goimports_output=$$($(GOIMPORTS) -d $$files 2>&1); \
			if [ "$$goimports_output" ]; then \
				echo "$$goimports_output"; \
				echo "goimports failed"; \
				echo "To fix it, run:"; \
				echo "goimports -w [FAILED_PACKAGE]"; \
				exit 1; \
			fi; \
		fi; \
	done
	@echo "goimports success"; \

goformat:
	@echo Checking if code is formatted
	@for package in $(PACKAGES); do \
		echo "Checking "$$package; \
		files=$$(go list -f '{{range .GoFiles}}{{$$.Dir}}/{{.}} {{end}}' $$package); \
		if [ "$$files" ]; then \
			gofmt_output=$$(gofmt -d -s $$files 2>&1); \
			if [ "$$gofmt_output" ]; then \
				echo "$$gofmt_output"; \
				echo "gofmt failed"; \
				echo "To fix it, run:"; \
				echo "go fmt [FAILED_PACKAGE]"; \
				exit 1; \
			fi; \
		fi; \
	done
	@echo "gofmt success"; \

check: govet lint goformat goimports
	@echo Checking for style guide compliance

docker:
	@if [ -z $$(docker ps -q -f name=mysql) ]; then \
		echo "Creating docker container..."; \
		docker rm -f mysql; \
		docker run -p 3306:3306 -e MYSQL_ROOT_PASSWORD=test -d --name mysql mysql:8.0 --default-authentication-plugin=mysql_native_password; \
		echo "Waiting for container to start..."; \
		until docker exec mysql mysqladmin ping -uroot -ptest --silent; do sleep 1; done; \
		echo "Creating database 'product'..."; \
		sleep 10; \
		docker exec mysql mysql -uroot -ptest -e "CREATE DATABASE IF NOT EXISTS product"; \
    else \
		echo "Docker container 'mysql' is already running."; \
    fi

godoc: $(GODOC)


prepare: dep check godoc

$(GOIMPORTS): ## Build goimports.
	GOBIN=$(TOOLS_BIN_DIR) $(GO_INSTALL) golang.org/x/tools/cmd/goimports $(GOIMPORTS_BIN) $(GOIMPORTS_VER)

$(GOLANGCILINT): ## Build golangci-lint
	GOBIN=$(TOOLS_BIN_DIR) $(GO_INSTALL) github.com/golangci/golangci-lint/cmd/golangci-lint $(GOLANGCILINT_BIN) $(GOLANGCILINT_VER)

$(GODOC): ## Build godoc.
	GOBIN=$(TOOLS_BIN_DIR) $(GO_INSTALL) github.com/princjef/gomarkdoc/cmd/gomarkdoc $(GODOC_BIN) $(GODOC_VER)