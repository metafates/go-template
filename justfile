#!/usr/bin/env just --justfile

go-mod := `go list`
flags := '-ldflags="-s -w"'

install:
    go install {{flags}}

build:
    go build {{flags}}

update:
    go get -u
    go mod tidy -v

rename new-go-mod:
    find . -type f -not -path './.git/*' -exec sed -i '' -e "s|{{go-mod}}|{{new-go-mod}}|g" {} \;