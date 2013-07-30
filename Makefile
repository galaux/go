all: build run

build:
	go build src/test.go

run:
	./test
	@echo -e '\n\nWaiting for source change'
