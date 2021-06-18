package main

import (
	"fmt"
)

type microsoft int

var x microsoft
var y int

func main() {
	fmt.Println("x:", x)
	fmt.Printf("%T\n", x)

	x = 42

	fmt.Println("x:", x)

	y = int(x)

	fmt.Println("y:", y)
	fmt.Printf("%T", y)
}