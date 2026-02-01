LITERATE_BIN := ./literate

all: generate test build

test:
	@echo "Running tests..."
	@go test ./...

build: generate
	@echo "Building..."
	@go build ./...

generate: $(LITERATE_BIN)
	@echo "Generating..."
	@$(LITERATE_BIN) README.md
	@echo "Formatting..."
	@go fmt ./...
	@echo "Tidying..."
	@go mod download
	@go mod tidy
	@echo "Vetting code..."
	@go vet ./...

$(LITERATE_BIN):
	@echo "Installing literate binary..."
	@GOBIN=$(shell pwd) go install github.com/schraf/literate@latest

clean:
	@echo "Cleaning..."
	@rm -f $(LITERATE_BIN)

help:
	@echo "Available targets:"
	@echo "  all      - Run generate, test and build"
	@echo "  generate - Generates the code from README.md"
	@echo "  test     - Run tests"
	@echo "  build    - Build the literate executable"
	@echo "  clean    - Deletes any intermediate files"
	@echo ""
	@echo "  help     - Show this help message"
