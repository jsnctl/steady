build:
	go build -o build ./cmd/steady/steady.go

test:
	go test ./...

all: build test
	./build