# Makefile for building the Go program

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=RegexGoler
BINARY_UNIX=$(BINARY_NAME)_unix
CMD_DIR=cmd

all: build

build:
	cd $(CMD_DIR) && $(GOBUILD) -o ../bin/$(BINARY_NAME) -v

clean:
	cd $(CMD_DIR) && $(GOCLEAN)
	rm -f bin/$(BINARY_NAME)
	rm -f bin/$(BINARY_UNIX)

run:
	cd $(CMD_DIR) && $(GOBUILD) -o /bin/$(BINARY_NAME) -v
	/bin/$(BINARY_NAME)
