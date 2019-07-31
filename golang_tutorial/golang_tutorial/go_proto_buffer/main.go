package main

//stat -f%z person.xml nie dziala :<

// go get github.com/golang/protobuf
//[!] błąd [!] can't load package: package github.com/golang/protobuf: no Go files in C:\Users\dp\go\src\github.com\golang\protobuf
//go get github.com/golang/protobuf/proto

//https://github.com/golang/protobuf
//go get -u github.com/golang/protobuf/protoc-gen-go dziala
//The compiler plugin, protoc-gen-go, will be installed in $GOBIN,
//defaulting to $GOPATH/bin. It must be in your $PATH for the protocol compiler, protoc, to find it.

//protoc
// 'protoc' is not recognized as an internal or external command,
// operable program or batch file.
import "fmt"

func main() {
	fmt.Println("Hello World")
}
