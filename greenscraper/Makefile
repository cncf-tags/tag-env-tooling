COMMIT_ID=$(shell git rev-parse --short HEAD)
VERSION=$(shell cat VERSION)
REPO=github.com/wrkode/greenscraper

APP_NAME=greenscraper

all: clean build

run:
	@go run main.go

clean:
	@echo ">> Cleaning..."
	@rm -rf bin

build: clean
	@echo ">> Building..."
	@mkdir bin
	@go build -v -o bin/$(APP_NAME) -ldflags "-X '$(REPO)/cmd.Version=$(VERSION) ($(COMMIT_ID))'" .

install: clean build
	@echo ">> Installing $(APP_NAME) in $(GOPATH)/bin..."
	@cp bin/$(APP_NAME) $(GOPATH)/bin

.PHONY: all clean build install run