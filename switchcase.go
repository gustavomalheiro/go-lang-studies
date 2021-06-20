package main

import (
	"fmt"
)

func main() {

	x := 1

	switch { // true por padrão
		case (x == 4), (x == 8):
			fmt.Println("é 1 ou 2")
		case (x < 10):
			fmt.Println("é 3 ou 4")
	}

	fmt.Println()
}