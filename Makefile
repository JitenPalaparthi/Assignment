hello:
	echo "Building and creating binary for Linux"

build: 
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/app -installsuffix cgo 