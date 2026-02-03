GOPATH := $(shell go env GOPATH)
GOBIN := $(GOPATH)/bin
LOCAL_BIN := $(GOBIN)/literate

PREFIX ?= /usr/local
BINDIR := $(PREFIX)/bin

.PHONY: all test build generate deps clean help install

all: generate test

test:
	@echo "Running tests..."
	@go test ./...

install: build
	@echo "Installing..."
	go install .
	mkdir -p $(DESTDIR)$(BINDIR)
	cp $(LOCAL_BIN) $(DESTDIR)$(BINDIR)/literate

build: generate
	@echo "Building..."
	@go build ./...

generate: main.go

main.go: README.md $(LOCAL_BIN)
	@echo "Generating..."
	@$(LOCAL_BIN) README.md
	@echo "Formatting..."
	@go fmt ./...
	@echo "Tidying..."
	@go mod download
	@go mod tidy
	@echo "Vetting code..."
	@go vet ./...

deps:
	@echo "Installing literate binary..."
	@./install.sh

$(LOCAL_BIN):
	@make deps

clean:
	@echo "Cleaning..."
	@rm -f literate *.go go.mod go.sum

help:
	@echo "Available targets:"
	@echo "  all      - Run generate, test and build"
	@echo "  generate - Generates the code from README.md"
	@echo "  test     - Run tests"
	@echo "  build    - Build the literate executable"
	@echo "  deps     - Install dependencies"
	@echo "  install  - Install literate build"
	@echo "  clean    - Deletes any intermediate files"
	@echo ""
	@echo "  help     - Show this help message"
