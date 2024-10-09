# Variables
BINARY_NAME := gchess_tui
SRC_DIR := ./cmd/$(BINARY_NAME)
BUILD_DIR := ./bin

# Default target
all: build

# Build the binary
build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_DIR)

# Run the application
run: build
	$(BUILD_DIR)/$(BINARY_NAME)

# Clean up
clean:
	rm -rf $(BUILD_DIR)

# Test the application
test:
	go test ./...

# Format the code
fmt:
	go fmt ./...

# Lint the code
lint:
	golangci-lint run

# Phony targets
.PHONY: all build run clean test fmt lint

