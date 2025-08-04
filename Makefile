GO := go

PKG := ./...

.PHONY: all build test clean help

all: build test

## Build the Go application
build:
	$(GO) build $(PKG)

## Run all tests
test:
	@echo "Running all tests..."
	$(GO) test  $(PKG) -v 

test-all:
	@echo "Running all tests..."
	$(GO) test  $(PKG) -v -bench=. -benchmem -cover

## Clean up build artifacts and test cache
clean:
	$(GO) clean
	$(GO) mod tidy -v
	rm -f $(shell find . -type f -name '*~' -o -name '#*#' -o -name '*.bak' -o -name 'tmp.*' -o -name 'core' )
	@echo "Cleaned up build artifacts and test cache."


