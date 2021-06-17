LINUX_AMD64 = CGO_ENABLED=0 GOOS=linux GOARCH=amd64

build:
	$(LINUX_AMD64) go build -o bexs api/main.go

linter:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$GOPATH/bin

lint:
	golangci-lint run ./...

start:
	FILE=storage/input-route.csv DEBUG=true HOST=0.0.0.0 PORT=3000 go run api/main.go

test:
	go test ./... -covermode=count -count 1 -v

dk-start:
	docker run -p 3000:3000 luizalabs:latest

dk-build: build-api
	docker build -t luizalabs:latest .

dk-deploy: dk-build dk-start
