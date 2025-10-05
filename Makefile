APP_NAME=ews-mailer

OUTPUT_DIR=./out
GIT_COMMIT=$(shell git rev-parse --short HEAD)
LDFLAGS=-ldflags "-s -w"

.PHONY: all build

all: build

build: main.go
	@echo "-- Build an application"
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${OUTPUT_DIR}/${APP_NAME}-${GIT_COMMIT}-linux-amd64 .
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ${OUTPUT_DIR}/${APP_NAME}-${GIT_COMMIT}-windows-amd64.exe .
