# Makefile

URL := "https://7260-206-190-232-194.ngrok.io"
REGISTRY_URL := quay.io
REGISTRY_USERNAME := techprober
VERSION := latest
ENV := prod
VERSION := latest
IMAGE_NAME := ci-bot

ifneq ($(ENV), dev)
	export IMAGE_TAG=$(VERSION)
else
	export IMAGE_TAG=dev
endif


.PHONY: login
login:
	@echo "==> Login to $(REGISTRY_URL)"
	@echo $(QUAY_PASS) | sudo nerdctl login $(REGISTRY_URL) -u $(REGISTRY_USERNAME) --password-stdin

.PHONY: build
build:
	@echo "==> Build application image with tag $(IMAGE_TAG)"
	@sudo nerdctl build \
		-t $(REGISTRY_URL)/$(REGISTRY_USERNAME)/$(IMAGE_NAME):$(IMAGE_TAG) \
		.

.PHONY: publish
publish: login build publish
	@echo "==> Push application image $(IMAGE_NAME):$(IMAGE_TAG) to $(REGISTRY_URL)"
	@sudo nerdctl push \
		$(REGISTRY_URL)/$(REGISTRY_USERNAME)/$(IMAGE_NAME):$(IMAGE_TAG) \

.PHONY: run
run:
	@echo "==> Run application locally"
	@sudo nerdctl run --rm -it \
		--name ci-bot \
		$(REGISTRY_URL)/$(REGISTRY_USERNAME)/$(IMAGE_NAME):$(IMAGE_TAG)

.PHONY: help
help:
	$(info ${HELP_MESSAGE})
	@exit 0

define HELP_MESSAGE

Usage: $ make [TARGETS]

TARGETS

	help            Show the help menu
	build           Build the application image
	run             Run the application container locally (VERSION optional)
	push            Build the application image, tag it with a custom version tag, and push it to the remote registry (Version required)

EXAMPLE USAGE

	build           Build the application image and tag it as latest
	run             Run the application container locally with the latest tag
	publish         Build the application iamge, tag it as v1.0.1, and push it to GHCR

endef
