OUTPUT_DIR=build
PROJECT_NAME := "nesgo"
TARGET=$(OUTPUT_DIR)/$(PROJECT_NAME)
PKG_LIST := $(shell go list ./... | grep -v /vendor)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor | grep -v _test.go)

.PHONY: all build clean test lint

all: build

lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}

test: ## Run unittests
	@gotest ${PKG_LIST} ./... -parallel 64

race: ## Run data race detector
	@go test -race -short ${PKG_LIST}

msan: ## Run memory sanitizer. Requires clang/llvm >= 3.8.0
	@go test -msan -short ${PKG_LIST}

build: ## Build the binary file
	@go build -o $(TARGET) -v

clean: ## Remove previous build
	@rm -f $(TARGET)

run: ## Run the build
	./$(TARGET)

