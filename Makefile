all: build

build:
	go build -ldflags "-linkmode external -extldflags -static" -o main
	chmod +x main

test:
	curl -L http://speedtest.ftp.otenet.gr/files/test1Mb.db -o test1Mb.db
	curl -L http://speedtest.ftp.otenet.gr/files/test100k.db -o test100k.db
	curl -X POST http://localhost:8081 -F file=@test1Mb.db -F key=test
	curl -X POST http://localhost:8081 -F file=@test100k.db -F key=test
	go clean -testcache ./...
	go test -v ./...

format:
	go fmt ./...
