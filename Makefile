NAME     := go-enjaxel
VERSION  := v0.0.1
REVISION := $(shell git rev-parse --short HEAD)

SRCS    := $(shell find . -type f -name '*.go')
LDFLAGS := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""

run:
	go run *.go

.PHONY: dep
dep:
	dep ensure

build: $(SRCS)
	go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o bin/$(NAME)

clean:
	rm -rf bin/*
	rm -rf vendor/*

.PHONY: test
test:
	go test

