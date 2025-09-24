# Makefile for the Greeting Example Application

# Variables
APP_NAME=greeting_example
BUILD_DIR=build

.DEFAULT_GOAL := help

.PHONY: fmt vet build run-server run-client clean help

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/server ./greeter_server/server.go
	go build -o $(BUILD_DIR)/client ./greeter_client/client.go

run-server:
	@echo "Starting gRPC server on port 50051..."
	go run ./greeter_server/server.go

run-client:
	@echo "Running gRPC client..."
	go run ./greeter_client/client.go

clean:
	rm -rf $(BUILD_DIR)
	go clean -cache


help:
	@echo "Makefile for the gRPC Greeting Example"
	@echo ""
	@echo "Usage:"
	@echo "  make fmt          # Format the code"
	@echo "  make vet          # Run go vet"
	@echo "  make build        # Build both server and client"
	@echo "  make run-server   # Run the gRPC server"
	@echo "  make run-client   # Run the gRPC client"
	@echo "  make clean        # Clean up build artifacts"
	@echo "  make help         # Show this help message"
	@echo ""
	@echo "To test the gRPC communication:"
	@echo "  1. In one terminal: make run-server"
	@echo "  2. In another terminal: make run-client"