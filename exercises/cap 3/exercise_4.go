package main

import (
	"fmt"
)

type microsoft int
var x microsoft

func main() {

	fmt.Println("x:", x)
	fmt.Printf("%T\n", x)

	x = 42

	fmt.Println("x:", x)
}