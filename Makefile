.PHONY: build

all: build container-build-run

build:
	go build -o ./release/starter-go cmd/starter-go/main.go

build-release:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s" -o ./release/starter-go cmd/starter-go/main.go
	strip ./release/starter-go

run:
	cd ./release && ./starter-go

test:
	ginkgo -randomizeAllSpecs -randomizeSuites -failOnPending -trace -race -progress -cover -r

# Container targets

container-build:
	docker build -f build/Containerfile -t vicenteherrera/starter-go .

container-run:
	docker run --rm -it -v "$$(pwd)"/test/in.yaml:/bin/in.yaml vicenteherrera/starter-go

container-build-run: container-build container-run
