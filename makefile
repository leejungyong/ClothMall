.PHONY: build clean tool lint help

all: build

build:
	go build -o ginbase cmd/main.go