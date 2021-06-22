BINARY=aquarium-lights

default: build

.PHONY: build
build:
	go fmt ./...
	go build -o ${BINARY}

.PHONY: release
release:
	go fmt ./...
	GOOS=darwin GOARCH=amd64 go build -o ./bin/${BINARY}_darwin_amd64
	GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY}_linux_amd64
	GOOS=linux GOARCH=arm64 go build -o ./bin/${BINARY}_linux_arm64
