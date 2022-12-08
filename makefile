
# Project variables -------------------------------------------

TARGET_BIN=starter-go
MAIN_DIR=.
GH_REPO=vicenteherrera/starter-go
CONTAINER_IMAGE=quay.io/vicenteherrera/starter-go


# Git repo info variables --------------------------------------

GIT_SHA := $(shell git -c log.showSignature=false rev-parse HEAD 2>/dev/null)
GIT_TAG := $(shell bash -c 'TAG=$$(git -c log.showSignature=false \
	describe --tags --exact-match --abbrev=0 $(GIT_SHA) 2>/dev/null); echo "$${TAG:-dev}"')

# Go compilation flags for automatic versioning on git tags and commit digest

LDFLAGS=-s -w \
        -X github.com/${GH_REPO}/cmd/${TARGET_BIN}.version=$(GIT_TAG) \
        -X github.com/${GH_REPO}/cmd/${TARGET_BIN}.commit=$(GIT_SHA) \
		-X github.com/${GH_REPO}/cmd/${TARGET_BIN}.date=$(date +"%Y-%m-%dT%H:%M:%S%z") \
		-X github.com/${GH_REPO}/cmd/${TARGET_BIN}.builtBy="makefile"

# Building the image -------------------------------------------------

.PHONY: all
# upgrade all dependencies, build binary and container, and run the container
all: upgrade build container-build container-run

.PHONY: upgrade
# upgrade dependencies
upgrade:
	go mod tidy

.PHONY: mod_download
# download dependencies
mod_download:
	go mod download

.PHONY: build
# build go binary
build:
	go build -ldflags="$(LDFLAGS)" -o ./release/${TARGET_BIN} ${MAIN_DIR}/main.go

.PHONY: build-release
# build go binary for release
build-release: mod_download vet test-noginkgo
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="$(LDFLAGS)" -o ./release/${TARGET_BIN} ${MAIN_DIR}/main.go
	strip ./release/${TARGET_BIN}

.PHONY: run
# run compiled go binary
run:
	cd ./release && ./${TARGET_BIN} --config ./config.yaml


# Lint ---------------------------------------------

.PHONY: lint
# execute all linters
lint: lint-go lint-yaml lint-containerfile

.PHONY: lint-go
# lint go
lint-go:
	golangci-lint run

.PHONY: lint-yaml
# lint yaml
lint-yaml:
	yamllint .

.PHONY: lint-containerfile
# lint containerfile
lint-containerfile:
	hadolint build/Containerfile


# Tests -------------------------------------------

.PHONY: test
# execute tests
test:
	ginkgo -randomize-all -randomize-suites -fail-on-pending -trace -race -progress -cover -r -v

 
# Dependencies ------------------------------------

.PHONY: 
# check all dependencies exist
dependencies:
	go version
	ginkgo version
	golangci-lint --version
	yamllint --version
	hadolint --version
	yaml --version

.PHONY: install_ginkgo
# install ginkgo
install_ginkgo:
	go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo
	go get github.com/onsi/gomega/...

.PHONY: install_golangci-lint
# install golangci-lint
install_golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.46.2

.PHONY: install_yamllint
# install yamllint
install_yamllint:
	pip install --user yamllint

.PHONY: install_yaml
# install yaml (pip)
install_yaml:
	pip install --user ruamel.yaml.cmd


# Container targets ------------------------------------------

# Check if sudo is required to run Docker
RUNSUDO := $(shell groups | grep ' docker ' 1>/dev/null || echo "sudo")

.PHONY: container-build
# build the container image
container-build:
	@echo "Building container image"
	@${RUNSUDO} docker build -f build/Containerfile -t ${CONTAINER_IMAGE} .

.PHONY: container-run
# run the container image
container-run:
	@echo "Running container image"
	@${RUNSUDO} docker run --rm -it \
		-v "$$(pwd)"/test/in.yaml:/bin/in.yaml \
		-u $$(id -u $${USER}):$$(id -g $${USER}) \
		${CONTAINER_IMAGE}

.PHONY: push
# push the container image
push:
	${RUNSUDO} docker push ${CONTAINER_IMAGE}

.PHONY: pull
# pull the container image
pull:
	${RUNSUDO} docker pull ${CONTAINER_IMAGE}
