MAKEFLAGS += --silent

ldflags := -s
ldflags += -w

build_flags := -ldflags=${ldflags}

go_mod := $(shell go list | awk -F/ '{print $$NF}')

all: help

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build        Build the ${go_mod} binary"
	@echo "  install      Install the ${go_mod} binary"
	@echo "  uninstall    Uninstall the ${go_mod} binary"
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
	@rm -f $(shell which ${go_mod})
