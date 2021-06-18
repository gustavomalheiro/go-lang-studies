package main

import (
	"fmt"
)

var x int = 10

func main() {

	// x = 20.5 (não pode, pois já foi declarado como integer)
	fmt.Printf("x: %v, %T", x, x)
}

