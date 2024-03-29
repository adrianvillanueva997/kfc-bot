run:
	go run ./src/main.go
run-dev:
	gow run ./src/main.go
build:
	go build -o app ./src
install:
	go mod download
docker-build:
	docker build -t kfcdiscordbot .
lint:
	golangci-lint run src/
fmt:
	go fmt ./...
sec:
	gosec ./...
docker-lint:
	hadolint --ignore DL3018 Dockerfile
docker-run-lint:
	docker run --rm -i hadolint/hadolint < Dockerfile