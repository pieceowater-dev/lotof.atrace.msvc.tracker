# Project configuration
APP_NAME = lotof.atrace.msvc.tracker
BUILD_DIR = bin
MAIN_FILE = cmd/server/main.go
PROTOC = protoc
PROTOC_GEN_GO = $(GOPATH)/bin/protoc-gen-go
PROTOC_GEN_GRPC_GO = $(GOPATH)/bin/protoc-gen-go-grpc
PROTOC_PKG = github.com/pieceowater-dev/lotof.atrace.proto
PROTOC_PKG_PATH = $(shell go list -m -f '{{.Dir}}' $(PROTOC_PKG))
PROTOC_DIR = protos
PROTOC_OUT_DIR = ./internal/core/grpc/generated

export PATH := /usr/local/bin:$(PATH)

.PHONY: all clean build run update setup grpc-gen grpc-clean grpc-update

# Default build target
all: build

# Setup the environment
setup: grpc-update
	@echo "Setup completed!"
	go mod tidy

# Update dependencies
update:
	go mod tidy

# Build the project
build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)

# Run the application
run: build
	./$(BUILD_DIR)/$(APP_NAME)

# Clean build artifacts
clean:
	rm -rf $(BUILD_DIR) gql-clean grpc-clean

# gRPC code generation
grpc-gen:
	@echo "Generating gRPC code from proto files..."
	mkdir -p $(PROTOC_OUT_DIR)
	find $(PROTOC_PKG_PATH)/$(PROTOC_DIR) -name "*.proto" | xargs $(PROTOC) \
		-I $(PROTOC_PKG_PATH)/$(PROTOC_DIR) \
		--go_out=$(PROTOC_OUT_DIR) \
		--go-grpc_out=$(PROTOC_OUT_DIR)
	@echo "gRPC code generation completed!"

# Clean gRPC generated files
grpc-clean:
	rm -rf $(PROTOC_OUT_DIR)

# Update gRPC dependencies
grpc-update:
	go get -u $(PROTOC_PKG)@latest
