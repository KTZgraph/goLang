//cwiczenie 1 str 23

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	start1 := time.Now()
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(i, " ", os.Args[i])
	}
	end1 := time.Now()
	fmt.Println("Koniec1", end1.Sub(start1))

	start2 := time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	end2 := time.Now()
	fmt.Println("Koniec2", end2.Sub(start2))

}
