BINARY_NAME=loggen

build:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux log-generator.go

container: build
	podman build .
run: build
	./${BINARY_NAME}

clean:
	go clean
