package main

import (
	"fmt"
)

func main() {
	x := 24
	y := x << 2
	z := x >> 2

	fmt.Printf("%b\n", y)
    fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", z)
}