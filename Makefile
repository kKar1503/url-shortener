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
