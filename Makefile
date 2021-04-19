prepare:
	go get ./...

lint:
	yamllint configs
	golangci-lint run --fix
	golangci-lint run --disable unused

test:
	cd pkg/order && godog
	go test -race -coverprofile=coverage.txt -covermode=atomic ./...

build:
	go build cmd/fdm/main.go

run:
	./main
	rm main

docker:
	docker-compose down
	docker-compose up --build

all: prepare lint test build run
