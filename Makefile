BINARY_NAME=SimpleMap

run:
	echo "Running ${BINARY_NAME}"
	go run ./cmd/simple-map

compile:
	echo "Compiling for every OS and Platform"
	GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin ./cmd/simple-map
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux ./cmd/simple-map
	GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}-windows ./cmd/simple-map

build:
	echo "Building ${BINARY_NAME} for local OS and Platform"
	go build -o bin/${BINARY_NAME} ./cmd/simple-map

clean:
	go clean
	rm -rf bin/

test:
	go test -cover -v ./...