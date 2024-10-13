.PHONY: run
.PHONY: dev
.PHONY: tidy
.PHONY: build
.PHONY: test
.PHONY: clean

run: build
	./bin/api

dev:
	go run ./cmd/api

tidy:
	go fmt ./...
	go mod tidy
	go mod verify

build:
	go build -ldflags="-s -w" -o=./bin/api ./cmd/api

test:
	go test ./...

clean:
	rm -rf bin
	rm -rf uploads
