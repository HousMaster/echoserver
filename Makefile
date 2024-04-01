.PHONY: all build run clean 

NAME=echoserver

all: build run

build:
	go build -o ./bin/$(NAME) ./cmd/$(NAME) 

run:
	./bin/$(NAME)

clean:
	rm ./bin/$(NAME)

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./bin/linux/$(NAME) ./cmd/$(NAME)
