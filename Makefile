install:
	go mod download
	go build .
	sudo mv ./odin /usr/local/bin
	odin --version

lint:
	golangci-lint run -E gofmt -E gci --fix;

build:
	go mod download
	mkdir -p bin/odin_darwin_amd64
	env GOOS=darwin GOARCH=amd64 go build -o bin/odin_darwin_amd64/odin
	mkdir -p bin/odin_darwin_arm64
	env GOOS=darwin GOARCH=arm64 go build -o bin/odin_darwin_arm64/odin
	mkdir -p bin/odin_linux_amd64
	env GOOS=linux GOARCH=amd64 go build -o bin/odin_linux_amd64/odin

compressed-builds: build
	tar -czvf bin/odin_darwin_amd64/odin_darwin_amd64.tar.gz bin/odin_darwin_amd64/odin
	tar -czvf bin/odin_darwin_arm64/odin_darwin_arm64.tar.gz bin/odin_darwin_arm64/odin
	tar -czvf bin/odin_linux_amd64/odin_linux_amd64.tar.gz bin/odin_linux_amd64/odin
