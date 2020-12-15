all: build

build:
	go build -ldflags "-linkmode external -extldflags -static" -o main
	chmod +x main

test:
	go clean -testcache ./...
	go test -v ./...

format:
	go fmt ./...
	