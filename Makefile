prepare:
	go get ./...

lint:
	yamllint configs
	golangci-lint run --fix --disable unused

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

local-deploy:
	minikube start
	kubectl apply -f backend-configmap.yaml
	kubectl apply -f backend-secret.yaml
	kubectl apply -f backend-deployment.yaml
	kubectl apply -f fdm-configmap.yaml
	kubectl apply -f fdm-deployment.yaml

all: prepare lint test build run
