install:
	go mod download
	go build .
	sudo mv ./odin /usr/local/bin
	odin --version

lint:
	golangci-lint run -E gofmt -E gci --fix;

build:
	go mod download
	env GOOS=darwin GOARCH=amd64 go build -o bin/odin_darwin_amd64
	# More OS and architectures can be supported -
	# env GOOS=linux GOARCH=amd64 go build -o bin/odin_linux_amd64
	# etc ...