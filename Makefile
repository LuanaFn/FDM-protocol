lint:
	yamllint configs
	golangci-lint run --disable unused

build:
	cd pkg/order && go mod vendor
	cd pkg/log && go mod vendor
	cd configs && go mod vendor
	cd cmd/open-marketplace && go mod vendor
	go mod vendor
	go build cmd/open-marketplace/main.go

test:
	cd pkg/order && godog

run:
	./main
	rm main

all: build lint test run
