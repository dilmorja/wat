# WebAssembly Test

Run a simple example of a server with static files and manipulate the DOM.

## Purpose

The purpose of this repository is to test and practice with WebAssembly on the server side and to be able to implement it in the ovni project. It also hopes to be of help to those who are just getting started on the subject.

## Requirements

- GNU Make
- Go

## Run

```bash
# Clone the repository
$ git clone https://github.com/dilmorja/wat.git
# Move to dir
$ cd wat
# Install dependencies
$ go mod tidy
# Compile app
$ make build
```

> **WARNING**: This way of compiling changes the environment variables. Before you build, make sure you know what the values are for your operating system and architecture. After compiling, set it to its default values: `go env -w GOOS=[your OS] GOARCH=[your arch]`.

## License

[MIT](LICENSE)