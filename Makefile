SHELL := /usr/bin/env bash

GIT_COMMIT=$(shell git rev-parse --verify HEAD)

GOOS = $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)
GOBUILD = go build -o bin/$(BINARY_BASENAME)-$(GOOS)-$(GOARCH) -ldflags "-X github.com/datawire/kubernaut/pkg/version.GitCommit=${GIT_COMMIT}"

BINARY_BASENAME=kubernaut

.PHONY: all build build.image build.image.devtools clean cloc fmt generate test.fast

all: clean fmt test.fast build

build: generate
	$(GOBUILD) main.go
	ln -sf $(BINARY_BASENAME)-$(GOOS)-$(GOARCH) bin/$(BINARY_BASENAME)

build.image:
	docker build \
	-t datawireio/kubernaut \
	-t datawireio/kubernaut:$(GIT_COMMIT) \
	-f Dockerfile \
	.

build.image.devtools:
	docker build \
	--build-arg UID=$(shell id -u) \
	-t knaut-dev \
	-f hack/docker/dev/Dockerfile \
	hack/docker/dev

clean:
	rm -rf bin

cloc: build.image.devtools
	docker run \
	--rm -it \
	--volume $(PWD):/project:ro \
	--workdir /project \
	knaut-dev \
	/usr/bin/cloc .

fmt:
	go fmt ./...

generate:
	protoc --go_out=plugins=grpc:. proto/bap.proto

test.fast:
	go test -tags=fast -v ./...
