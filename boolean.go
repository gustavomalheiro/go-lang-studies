package main

import (
	"fmt"
)

var x bool

func main() {
	fmt.Println(x) // zero value
	x = true
	fmt.Println(x) // valor atribuido

	x = 10 < 100 // true
	y := (10 > 100) // false
	z := (10 == 100) // false

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}