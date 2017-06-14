#
# Makefile
# Common and useful commands for PosInsertOrder
#

.PHONY: run, build, fmt
default: run

run:
	@go run main.go

build:
	@go build

fmt:
	@gofmt -w -l .

test:
	@go test -v ./...
