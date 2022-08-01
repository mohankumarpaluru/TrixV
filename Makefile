-include environ.inc
.PHONY: help deps dev build install image release test clean clean-all

export CGO_ENABLED=0
VERSION=$(shell git describe --abbrev=0 --tags 2>/dev/null || echo "0.0.0")
COMMIT=$(shell git rev-parse --short HEAD || echo "HEAD")
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
GOCMD=go
GOVER=$(shell go version | grep -o -E 'go1\.17\.[0-9]+')

DESTDIR=/usr/local/bin

ifeq ($(BRANCH), master)
IMAGE := prologic/tube
TAG := latest
else
IMAGE := prologic/tube
TAG := dev
endif

all: help

.PHONY: help

help: ## Show this help message
	@echo "tube - a Youtube-like (without censorship and features you don't need!) Video Sharing App"
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

preflight: ## Run preflight checks to ensure you have the right build tools
	@./preflight.sh

deps: ## Install any dependencies required

dev : DEBUG=1
dev : server ## Build debug version of tube
	@./tube

server: generate ## Build the tube server
	@$(GOCMD) build -tags "embed netgo static_build" -installsuffix netgo \
		-ldflags "-w \
		-X $(shell go list).Version=$(VERSION) \
		-X $(shell go list).Commit=$(COMMIT)" \
		.

build: server ## Build the server

generate: ## Genereate any code required by the build
	@if [ x"$(DEBUG)" = x"1"  ]; then		\
	  echo 'Running in debug mode...';	\
	fi

install: build ## Install tube to $DESTDIR
	@install -D -m 755 tube $(DESTDIR)/tube

ifeq ($(PUBLISH), 1)
image: generate ## Build the Docker image
	@docker build --build-arg VERSION="$(VERSION)" --build-arg COMMIT="$(COMMIT)" -t $(IMAGE):$(TAG) .
	@docker push $(IMAGE):$(TAG)
else
image: generate
	@docker build --build-arg VERSION="$(VERSION)" --build-arg COMMIT="$(COMMIT)" -t $(IMAGE):$(TAG) .
endif

release: generate ## Release a new version to Gitea
	@./tools/release.sh

fmt: ## Format sources fiels
	@$(GOCMD) fmt ./...

test: ## Run test suite
	@CGO_ENABLED=1 $(GOCMD) test -v -cover -race ./...

coverage: ## Get test coverage report
	@CGO_ENABLED=1 $(GOCMD) test -v -cover -race -cover -coverprofile=coverage.out  ./...
	@$(GOCMD) tool cover -html=coverage.out

clean: ## Remove untracked files
	@git clean -f -d -x

clean-all:  ## Remove untracked and Git ignores files
	@git clean -f -d -X
