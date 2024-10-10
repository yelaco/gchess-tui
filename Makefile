# Variables
BINARY_NAME := gchess_tui
SRC_DIR := ./cmd/$(BINARY_NAME)
BUILD_DIR := ./bin

# Default target
all: build

# Build the binary
build:
	@if [ ! -f .env ]; then cp .env.example .env; fi
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_DIR)

# Run the application
run: build
	@if ls *.log 1> /dev/null 2>&1; then rm *.log; fi
	$(BUILD_DIR)/$(BINARY_NAME)

# Debug the application
debug: export debug=true
debug: run

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

