//testowanie w GO
//tutaj kod do przetestowania
//uwaga jak w wielu plikach jest package main, to test sie wywali na twarz
//pliki umieszcza wosobnych folderach jak tutaj
// λ go test
// # _/C_/Users/dp/Desktop/PracowaniaProgramowania/backend [_/C_/Users/dp/Desktop/PracowaniaProgramowania/backend.test]
// .\duplication.go:10:6: main redeclared in this block
//         previous declaration at .\cw1.go:12:6
// .\echo1.go:8:6: main redeclared in this block
//         previous declaration at .\duplication.go:10:6
// .\echo2.go:9:6: main redeclared in this block
//         previous declaration at .\echo1.go:8:6
// .\hello.go:5:6: main redeclared in this block
//         previous declaration at .\echo2.go:9:6
// .\serverTest.go:12:6: main redeclared in this block
//         previous declaration at .\hello.go:5:6
// .\serverTest2.go:15:6: main redeclared in this block
//         previous declaration at .\serverTest.go:12:6
// .\tut1.go:37:6: main redeclared in this block
//         previous declaration at .\serverTest2.go:15:6
// .\tut2.go:13:6: Article redeclared in this block
//         previous declaration at .\tut1.go:11:6
// .\tut2.go:19:6: Articles redeclared in this block
//         previous declaration at .\tut1.go:17:6
// .\tut2.go:21:6: allArticles redeclared in this block
//         previous declaration at .\tut1.go:19:44
// .\tut2.go:21:6: too many errors
// FAIL    _/C_/Users/dp/Desktop/PracowaniaProgramowania/backend [build failed]
package main

import (
	"fmt"
)

func Calculate(x int) (result int) {
	result = x + 2
	return result
}

func main() {
	fmt.Println("Go testing tutorial")
	result := Calculate(2)
	fmt.Println(result)
}
