
build:
	go build cmd/

run:
	./starter-go

test:
	ginkgo -randomizeAllSpecs -randomizeSuites -failOnPending -trace -race -progress -cover -r

container-build:
	docker build -f build/Containerfile -t vicenteherrera/starter-go .

container-run:
	docker run --rm -it vicenteherrera/starter-go --help

container-build-run: container-build container-run
