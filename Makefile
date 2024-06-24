# Makefile for node-prometheus

# Variables
IMAGE  = cyrip/go-prometheus
TAG = 1.0.0

# targets
.PHONY: all

all: build

# Target to run the Ansible playbook
build:
	docker buildx build -t ${IMAGE}:${TAG} .
	docker push ${IMAGE}:${TAG}
run:
	docker run --rm --init -p 8080:8080 ${IMAGE}:${TAG}