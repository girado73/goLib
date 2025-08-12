# Project name (binary output name)
BINARY_NAME := goLib

# Directories
BIN_DIR := bin

# Default target
all: build

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN_DIR)/$(BINARY_NAME) .

# Run the binary
run: build
	@./$(BIN_DIR)/$(BINARY_NAME)

# Clean up build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(BIN_DIR)

# Install dependencies
deps:
	@go mod tidy

.PHONY: all build run clean deps
