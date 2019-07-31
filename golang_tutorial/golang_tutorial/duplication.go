// wyswietlanie tekstu kazdej linii która pojawiła się na IO
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
		// fmt.Println(counts)
	}
	// //UWAGA: ignorowowanie potencjalnych błędów z funkcji input.Err()
	for line, n := range counts {
		// z jakiegos powodu
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			// fmt.Printf("%d", n)
		}
	}

}
