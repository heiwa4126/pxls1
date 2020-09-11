#!/bin/sh -ex
cd $(dirname $0)/..
BIN=$(basename $PWD)

go mod tidy
go fmt ./...

# test
go clean -testcache
go test ./...

# lint
go vet ./...
shadow ./...
golangci-lint run

# build
VER=$(git tag --sort=-v:refname | grep '^v' | head -1)
REV=$(git rev-parse --short HEAD)
go build -ldflags "-s -w -X main.Version=$VER -X main.Revision=$REV" -trimpath

# omake
go version -m "$BIN"
upx "$BIN"
