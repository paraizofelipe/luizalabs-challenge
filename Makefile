LINUX_AMD64 = CGO_ENABLED=0 GOOS=linux GOARCH=amd64

deps:
	go mod tidy
	go mod download

build:
	$(LINUX_AMD64) go build -o luizalabs-challenge main.go

linter:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$GOPATH/bin

lint:
	golangci-lint run ./...

mockgen:
	@go get github.com/golang/mock/mockgen@v1.5.0
	mockgen -source ./buyer/repository/repository.go -destination ./buyer/repository/repository_mock.go -package repository
	mockgen -source ./product/repository/repository.go -destination ./product/repository/repository_mock.go -package repository
	mockgen -source ./buyer/service/service.go -destination ./buyer/service/service_mock.go -package service
	mockgen -source ./product/service/service.go -destination ./product/service/service_mock.go -package service

start: build
	go run main.go

test:
	go test ./... -covermode=count -count 1

dk-start:
	docker run -p 3000:3000 luizalabs:latest

dk-build: build
	docker build -t luizalabs:latest .

dk-deploy:
	docker-compose up -d --build
