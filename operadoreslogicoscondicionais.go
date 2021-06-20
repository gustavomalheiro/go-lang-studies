package main

import (
	"fmt"
)

func main() {

	x := 4

	if x == 2 || x == 3 {
		fmt.Println("xis é dois ou três")
	} else if x % 2 == 0 && x == 4 {
		fmt.Println("é 4")
	}

	// desafio

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}