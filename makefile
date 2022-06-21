.PHONY: build test

TARGET_BIN=starter-go
MAIN_DIR=cmd/starter-go
CONTAINER_IMAGE=vicenteherrera/starter-go

all: build container-build-run

build:
	go build -o ./release/${TARGET_BIN} ${MAIN_DIR}/main.go

build-release:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s" -o ./release/${TARGET_BIN} ${MAIN_DIR}/main.go
	strip ./release/${TARGET_BIN}

run:
	cd ./release && ./${TARGET_BIN}

test:
	ginkgo -randomize-all -randomize-suites -fail-on-pending -trace -race -progress -cover -r

update:
	go mod tidy

# dependencies

dependecies:
	go version
	ginkgo version

install_ginkgo:
	go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo
	go get github.com/onsi/gomega/...

# Container targets

container-build:
	@if groups $$USER | grep -q '\bdocker\b'; then RUNSUDO="" ; else RUNSUDO="sudo" ; fi && \
	    $$RUNSUDO docker build -f build/Containerfile -t ${CONTAINER_IMAGE} .

container-run:
	@if groups $$USER | grep -q '\bdocker\b'; then RUNSUDO="" ; else RUNSUDO="sudo" ; fi && \
	    $$RUNSUDO docker run --rm -it \
		-v "$$(pwd)"/test/in.yaml:/bin/in.yaml \
		-u $$(id -u $${USER}):$$(id -g $${USER}) \
		${CONTAINER_IMAGE}

container-build-run: container-build container-run
