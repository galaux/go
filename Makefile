all: build run

build:
	go build src/test.go

run:
	./test
	@echo -e '\nWaiting for source change'
