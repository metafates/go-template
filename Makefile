MAKEFLAGS += --silent

ldflags := -s 
ldflags += -w

build_flags := -ldflags=${ldflags}

app := $(shell grep -m 1 go go.mod | cut -d\  -f2 | awk -F/ '{print $$NF}')
go_version_required := $(shell grep -m 2 go go.mod | tail -n 1 | cut -d\  -f2)

# Checks if the go compiler is installed and if it is the correct version
define check_go
	@if [ -z "$$(which go)" ]; then \
		echo "go is not installed"; \
		exit 1; \
	else \
		if [ "$$(go version | cut -d\  -f3 | sed 's/go//g')" \< "${go_version_required}" ]; then \
			echo "go version is less than required version"; \
			exit 1; \
		fi; \
	fi
endef

define print_yellow
	@echo "\033[1;33m$1\033[0m"
endef

define print_faint
	@echo "\033[2m$1\033[0m"
endef

define print_green
	@echo "\033[1;32m$1\033[0m"
endef

all: help

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build        Build the ${app} binary"
	@echo "  install      Install the ${app} binary"
	@echo "  uninstall    Uninstall the ${app} binary"
	@echo "  test         Run the tests"
	@echo "  help         Show this help message"
	@echo ""

install:
	$(call check_go)
	$(call print_faint,"Installing ${app}...")
	@go install "$(build_flags)"
	$(call print_green,Installed)


build:
	$(call check_go)
	$(call print_faint,"Building ${app}...")
	@go build "$(build_flags)"
	$(call print_green,Built)

test:
	$(call check_go)
	@go test ./...

uninstall:
	@rm -f $(shell which ${app})
	$(call print_yellow,Uninstalled)

.PHONY: all help install build test uninstall
