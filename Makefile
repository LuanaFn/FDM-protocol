install:
	go mod tidy
	go get -u ./...

lint:
	yamllint configs .github/workflows
	golangci-lint run --fix --disable unused
	gofmt -l -w -s .
	golint ./...

test:
	go test -v ./...

build:
	go build cmd/fdm/main.go

run:
	./main
	rm main

docker:
	docker-compose -f deploy/docker-compose.yml down
	docker-compose -f deploy/docker-compose.yml up --build

local-deploy:
	minikube start
	kubectl apply -f deploy/backend-configmap.yaml
	kubectl apply -f deploy/backend-secret.yaml
	kubectl apply -f deploy/backend-deployment.yaml
	kubectl apply -f deploy/fdm-configmap.yaml
	kubectl apply -f deploy/fdm-deployment.yaml

deploy:
	cd deploy && terraform apply

all: install lint test docker
