
build:
	go build cmd/

run:
	./starter-go


container-build:
	docker build -f build/Containerfile -t vicenteherrera/starter-go .

container-run:
	docker run --rm -it vicenteherrera/starter-go --help

container-build-run: container-build container-run
