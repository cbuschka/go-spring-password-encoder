PROJECT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
ifeq (${GOPATH},)
	GOPATH := ${HOME}/go
endif

all:	build

lint:
	@echo "### Running linter..."
	go get -u golang.org/x/lint/golint
	${GOPATH}/bin/golint ./...

build:	test lint
	go vet ./...

format:
	gofmt -l -w -s ${PROJECT_DIR}

test:
	@echo "### Running unit tests..."
	go test -cover -race -coverprofile=coverage.txt -covermode=atomic ./...
