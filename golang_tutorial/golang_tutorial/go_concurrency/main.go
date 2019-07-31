package main

import (
	"fmt"
	"time"
)

func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("Concurrency With GoRoutines")

	go compute(5) //asynchroniczne wywolania goutis routine
	go compute(5)

	fmt.Scanln() //zeby petla glowa(main) nie zakoncyl sie przed watkami
}
