install:
	go build .
	sudo mv ./odin /usr/local/bin
	odin --version

lint:
	find ./ -type f -name "*.go" -exec go fmt "{}" \;
