build:
	@go build -o ./wat.exe ./app.go
	@go env -w GOOS=js GOARCH=wasm
	@go build -o ./bin/go.wasm ./wasm.go