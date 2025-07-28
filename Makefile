BUF_VERSION := v1.55.1
PROTOC_GEN_VERSION := v1.31.0
GOLANGCI_LINT_VERSION := v2.3.0
GOBIN := $(shell go env GOPATH)/bin

.PHONY: default
default: help

# Self-documenting Makefile
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: setup
setup: ## Install project dependencies
	go install github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION)
	go install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_VERSION)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh \
		| sh -s -- -b $(GOBIN) $(GOLANGCI_LINT_VERSION)

.PHONY: proto
proto: ## Generate Go code from protobuf definitions using buf
	PATH=$(GOBIN) buf generate

.PHONY: build
build: proto ## Build the project
	go build -o bin/weather cmd/weather/main.go

.PHONY: run
run: ## Run the server
	go run cmd/weather/main.go

.PHONY: test
test: ## Run tests
	go test ./...

.PHONY: test-coverage
test-coverage: ## Run tests with code coverage
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

.PHONY: lint
lint: ## Run golangci-lint
	$(GOBIN)/golangci-lint run

.PHONY: clean
clean: ## Clean generated files
	rm -rf bin/
	rm -f coverage.out coverage.html
