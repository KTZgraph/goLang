package main

//[!] DOESNT WORK ON WINDOWS [!]
//backend dla frontu webowego
//GOARCH=wasm GOOS=js go build -o lib.wasm main.go
// go env
// set GOARCH=wasm
// set GOOS=js
// go build -o lib.wasm main.go

func main() {
	println("Hello World")
}
