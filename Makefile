install:
	go mod download
	go build .
	sudo mv ./odin /usr/local/bin
	odin --version

lint:
	golangci-lint run -E gofmt -E gci --fix;
