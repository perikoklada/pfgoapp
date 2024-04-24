.DEFAULT_GOAL := goapp

.PHONY: all
all: clean goapp client

.PHONY: goapp
goapp:
	mkdir -p bin
	go build -o bin ./cmd/server/...

.PHONY: client
client:
	mkdir -p bin
	go build -o bin ./cmd/client/...

.PHONY: clean
clean:
	go clean
	rm -f bin/*
