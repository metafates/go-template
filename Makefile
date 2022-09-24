MAKEFLAGS += --silent

ldflags := -s
ldflags += -w

build_flags := -ldflags=${ldflags}

all: help

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build        Build the binary"
	@echo "  install      Install the binary"
	@echo "  uninstall    Uninstall the binary"
	@echo "  test         Run the tests"
	@echo "  help         Show this help message"
	@echo ""

install:
	@go install "$(build_flags)"


build:
	@go build "$(build_flags)"

test:
	@go test ./...

uninstall:
	@rm -f $(shell which teleq)
