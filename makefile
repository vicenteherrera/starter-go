.PHONY: build test

all: build container-build-run

build:
	go build -o ./release/starter-go cmd/starter-go/main.go

build-release:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s" -o ./release/starter-go cmd/starter-go/main.go
	strip ./release/starter-go

run:
	cd ./release && ./starter-go

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
	    $$RUNSUDO docker build -f build/Containerfile -t vicenteherrera/starter-go .

container-run:
	@if groups $$USER | grep -q '\bdocker\b'; then RUNSUDO="" ; else RUNSUDO="sudo" ; fi && \
	    $$RUNSUDO docker run --rm -it -v "$$(pwd)"/test/in.yaml:/bin/in.yaml vicenteherrera/starter-go

container-build-run: container-build container-run
