.PHONY: help build install clean test lint run example

BINARY_NAME=persiancal
MAIN_PATH=./cmd/persiancal
INSTALL_PATH=$(shell go env GOPATH)/bin

help:
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'

build: ## Build the CLI binary
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) $(MAIN_PATH)
	@echo "Build complete: ./$(BINARY_NAME)"

install: ## Install the CLI to GOPATH/bin
	@echo "Installing $(BINARY_NAME) to $(INSTALL_PATH)..."
	@go install $(MAIN_PATH)
	@echo "Installed to $(INSTALL_PATH)/$(BINARY_NAME)"

clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)
	@rm -rf dist/
	@go clean
	@echo "Clean complete"

test:
	@echo "Running tests..."
	@go test -v ./...

test-coverage:
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.txt -covermode=atomic ./...
	@go tool cover -html=coverage.txt -o coverage.html
	@echo "Coverage report generated: coverage.html"

lint:
	@echo "Running linters..."
	@go vet ./...
	@go fmt ./...
	@echo "Lint complete"

run: build
	@./$(BINARY_NAME)

example:
	@echo "Running example..."
	@go run example/main.go

mod-tidy:
	@echo "Running go mod tidy..."
	@go mod tidy

mod-download:
	@echo "Downloading dependencies..."
	@go mod download

check: lint test

all: clean lint test build

.DEFAULT_GOAL := help

