SHELL := /usr/bin/env bash

build.systemd.img:
	docker build \
	-t knaut-systemd \
	-f docker/systemd/Dockerfile \
	.

run.systemd.shell: build.systemd.img
	docker run \
	--rm \
	-it \
	--security-opt seccomp=unconfined \
	--tmpfs /run \
	--tmpfs /run/lock \
	--volume /sys/fs/cgroup:/sys/fs/cgroup:ro \
	knaut-systemd

build.dev.img:
	docker build \
	--build-arg UID=$(shell id -u) \
	-t knaut-dev \
	-f docker/dev/Dockerfile \
	.

run.dev.shell: build.dev.img
	docker run \
	--rm \
	-it \
	--volume $(PWD):/hack \
	--workdir /hack \
	knaut-dev \
	/bin/bash
