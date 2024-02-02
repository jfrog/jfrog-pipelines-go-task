SHELL := /bin/bash
GOCMD = go
install:
	go mod download

tidy:
	go mod tidy

run:
	go run main.go $(ARGS)

test:
	go test -v --cover ./...

build-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build $(ARGS) -o bin/{{ .name }}-Linux-x86_64
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build $(ARGS) -o bin/{{ .name }}-linux-x86_64

build-mac:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build $(ARGS)

build-all:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build $(ARGS) -o bin/{{ .name }}-Linux-x86_64
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build $(ARGS) -o bin/{{ .name }}-linux-x86_64
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build $(ARGS) -o bin/{{ .name }}-Linux-ARM64
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build $(ARGS) -o bin/{{ .name }}-Darwin-x86_64
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build $(ARGS) -o bin/{{ .name }}-Darwin-ARM64
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build $(ARGS) -o bin/{{ .name }}-Windows-x86_64

build-all-ci:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build $(ARGS) -o bin/{{ .name }}-Linux-x86_64 -ldflags "-s -w"
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build $(ARGS) -o bin/{{ .name }}-Linux-ARM64 -ldflags "-s -w"
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build $(ARGS) -o bin/{{ .name }}-Darwin-x86_64 -ldflags "-s -w"
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build $(ARGS) -o bin/{{ .name }}-Darwin-ARM64 -ldflags "-s -w"
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build $(ARGS) -o bin/{{ .name }}-Windows-x86_64 -ldflags "-s -w"

do-all: install tidy test build-all-ci
