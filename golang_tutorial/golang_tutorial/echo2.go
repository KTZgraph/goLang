//Echo2 wyświetla swoje argumenty wiersza poleceń
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args { //taki foreach
		fmt.Println(arg)
		s += sep + arg //do zmiennej arg przypisana jest lista
		sep = " "
	}
	fmt.Println(s)
	// fmt.Println(os.Args)
	// fmt.Println(strings.Join(os.Args, " "))

}
