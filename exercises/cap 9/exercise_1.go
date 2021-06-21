package main

import (
	"fmt"
)

func main() {

	num := [5]int{12, 21, 34, 412, 512}

	for i, v := range num {
		fmt.Println("Índice:", i, "Valor:", v)
	}

	fmt.Printf("%T", num)
}