all: deps generate test build

test:
	@echo "Running tests..."
	@go test ./...

build:
	@echo "Building..."
	@go build ./...

generate:
	@echo "Generating..."
	@./cmd README.md
	@echo "Formatting..."
	@go fmt ./...
	@echo "Tidying..."
	@go mod download
	@go mod tidy
	@echo "Vetting code..."
	@go vet ./...

deps:
	@echo "Installing dependencies..."
	@GOBIN=$(shell pwd) go install github.com/schraf/literate/cmd@latest

clean:
	@echo "Cleaning..."
	@rm -f literate

help:
	@echo "Available targets:"
	@echo "  all    - Run generate, test and build"
	@echo "  test   - Run tests"
	@echo "  build  - Build the literate executable"
	@echo "  deps   - Install dependencies"
	@echo "  clean  - Deletes any intermediate files"
	@echo ""
	@echo "  help   - Show this help message"
