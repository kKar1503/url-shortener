run:
	go run cmd/server/main.go

build:
	go build -o bin/main cmd/server/main.go

tidy:
	go mod tidy

download:
	go mod download

test:
	go test -v ./...

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux main.go

build-docker:
	docker build -t url-shortener .

run-docker:
	docker run -p 8080:8080 url-shortener
