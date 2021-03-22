prepare:
	cd pkg/order && go mod vendor
	cd pkg/log && go mod vendor
	cd configs && go mod vendor
	cd cmd/open-marketplace && go mod vendor
	go mod vendor

lint:
	yamllint configs
	golangci-lint run --disable unused

test:
	cd pkg/order && godog

build:
	go build cmd/open-marketplace/main.go

run:
	./main
	rm main

all: prepare lint test build run
