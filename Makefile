# Makefile for a Go project

# Binary name
BINARY_NAME=oapi-demo

# Build the binary
build:
	go build -o ${BINARY_NAME} main.go

# Run the program
run: build
	./${BINARY_NAME}

# Clean up
clean:
	go clean
	rm ${BINARY_NAME}

# Fetch dependencies (useful if you're using external packages)
deps:
	go mod tidy