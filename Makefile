SHELL := /usr/bin/env bash

GIT_COMMIT=$(shell git rev-parse --verify HEAD)

GOOS = $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)
GOBUILD = go build -o bin/$(BINARY_BASENAME)-$(GOOS)-$(GOARCH) -ldflags "-X github.com/datawire/kubernaut/pkg/version.GitCommit=${GIT_COMMIT}"

BINARY_BASENAME=kubernaut

all: clean build

clean:
	rm -rf bin

build:
	$(GOBUILD) main.go
	ln -sf $(BINARY_BASENAME)-$(GOOS)-$(GOARCH) bin/$(BINARY_BASENAME)

build.image:
	docker build \
	-t datawireio/kubernaut \
	-t datawireio/kubernaut:$(GIT_COMMIT) \
	-f Dockerfile \
	.

test.fast:
	go test -tags=fast -v ./...
