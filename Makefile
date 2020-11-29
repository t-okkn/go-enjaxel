NAME     := go-enjaxel
VERSION  := v0.0.2
REVISION := $(shell git rev-parse --short HEAD)

SRCS    := $(shell find . -type f -name '*.go')
LDFLAGS := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""

GOVER     := $(shell go version | awk '{ print substr($$3, 3) }' | tr "." " ")
VER_JUDGE := $(shell if [ $(word 1,$(GOVER)) -eq 1 ] && [ $(word 2,$(GOVER)) -le 10 ]; then echo 0; else echo 1; fi)

run:
	go run *.go

init:
ifeq ($(VER_JUDGE),1)
	go mod init
else
	echo "Packageの取得は手動で行ってください"
endif

build: $(SRCS)
	go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o bin/$(NAME)

clean:
	rm -rf bin/*
	rm -rf vendor/*

.PHONY: test
test:
	go test

