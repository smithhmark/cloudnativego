# Project variables
BINARY_NAME=myapp
GO_FILES=$(shell find . -name "*.go")

# Default target: build the binary
all: build

# Build the binary
build-libs:
	@echo "building internal..."
	@go build ./internal/...

CMDS = $(shell ls ./cmd)
build: build-libs
	@echo "building commands:"
	@for cmd in $(CMDS); do \
		echo "-> $$cmd"; \
		#echo go build -o bin/$$cmd ./$$cmd/...; \
		go build -o bin/$$cmd ./cmd/$$cmd/...; \
	done

# Run the application
run: build
	./bin/$(bin)

# Run tests
test:
	go test ./...

tidy:
	go mod tidy

# Run tests with coverage
coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Lint the code (requires golangci-lint)
lint:
	golangci-lint run

# Clean build artifacts
clean:
	go clean
	rm -rf bin/
	rm -f coverage.out

# Help command to list targets
addval:
	curl -X PUT -d 'Hello, k-v store!' -v http://localhost:8080/v1/key/key-a

getval:
	curl -X GET  -v http://localhost:8080/v1/key/key-a

delval:
	curl -X DELETE  -v http://localhost:8080/v1/key/key-a
help:
	@echo "Available targets:"
	@echo "  make build-libs         - Build all binaries"
	@echo "  make build              - Build all binaries"
	@echo "  make run bin=<bin-name> - Build and run the binary"
	@echo "  make test               - Run tests"
	@echo "  make coverage           - Run tests and show coverage"
	@echo "  make lint               - Run linter"
	@echo "  make clean              - Remove build artifacts"

.PHONY: all build build-libs run test coverage lint clean help
