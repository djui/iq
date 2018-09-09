.PHONY: build
build:
	go build

.PHONY: install
install:
	go install

.PHONY: release
release:
	GOOS=darwin GOARCH=amd64 go build -o iq-darwin-amd64 .
	GOOS=linux GOARCH=amd64 go build -o iq-linux-amd64 .
