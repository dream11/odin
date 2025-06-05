install:
	protoc --proto_path=proto/ --go_out=. --go_opt=module=github.com/dream11/odin --go-grpc_out=. --go-grpc_opt=module=github.com/dream11/odin $(shell find proto -name '*.proto')
	go mod download
	go build .
	sudo mv ./odin /usr/local/bin
	odin version

lint:
	golangci-lint run -E gofmt -E gci --fix;

build:
	go mod download
	mkdir -p bin/odin_darwin_amd64
	env GOOS=darwin GOARCH=amd64 go build -o bin/odin_darwin_amd64/odin
	mkdir -p bin/odin_darwin_arm64
	env GOOS=darwin GOARCH=arm64 go build -o bin/odin_darwin_arm64/odin
	mkdir -p bin/odin_linux_amd64
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/odin_linux_amd64/odin

compressed-builds: build
	cd bin/odin_darwin_amd64 && tar -czvf ../odin_darwin_amd64.tar.gz odin
	cd bin/odin_darwin_arm64 && tar -czvf ../odin_darwin_arm64.tar.gz odin
	cd bin/odin_linux_amd64 && tar -czvf ../odin_linux_amd64.tar.gz odin
