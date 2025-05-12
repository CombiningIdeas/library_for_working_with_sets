# Path to Go
GO ?= go

# Application name
APP_NAME := libraryForWorkingWithSets

# Build directory
BUILD_DIR := build
PACKAGE_DIR := $(BUILD_DIR)/package

# Default target
all: build

#Building an executable file
build:
	mkdir -p $(BUILD_DIR)/package
	$(GO) build -v -o $(PACKAGE_DIR)/$(APP_NAME) ./cmd/$(APP_NAME)

# Starting the application
run:
	$(GO) run ./cmd/$(APP_NAME)

# Formatting code
fmt:
	$(GO) fmt ./...

# Code review
vet:
	$(GO) vet ./...

# Testing
test:
	$(GO) test ./...

# Assembly Cleanup
clean:
	rm -rf $(BUILD_DIR)

# Installing dependencies
deps:
	$(GO) mod tidy

# Complete rebuild
rebuild: clean deps build
