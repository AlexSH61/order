SHELL := /bin/bash
KIND := kindest/node:v1.29.1
KIND_CLUSTER := shop-cluster
BASE_IMAGE_NAME := eshop
SERVICE_NAME := order
VERSION := 0.0.1
SERVICE_IMAGE := $(BASE_IMAGE_NAME)/$(SERVICE_NAME):$(VERSION)

img:
	docker build \
		-f infra/docker/order.Dockerfile \
		-t $(SERVICE_IMAGE) \
		.

dev-up:
	kind create cluster \
	--image $(KIND) \
	--name $(KIND_CLUSTER) \
	--config infra/k8s/dev/kind_config.yaml
dev-down:
	kind delete cluster --name $(KIND_CLUSTER)
dev-load:
    kind load docker-image $(SERVICE_IMAGE) --name $(KIND_CLUSTER)