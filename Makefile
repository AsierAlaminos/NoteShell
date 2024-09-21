BINARY_NAME=noteShell

build:
	go build -o ${BINARY_NAME} ./cmd/app

run: build
	./${BINARY_NAME}

clean:
	rm -rf ${BINARY_NAME}

.PHONY: build run clean
all: build
