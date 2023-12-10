build:
	go build -o steady-binary ./cmd/steady/steady.go

test:
	go test ./...

all: build test
	./steady-binary