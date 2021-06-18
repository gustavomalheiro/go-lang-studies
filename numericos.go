package main

import (
	"fmt"
	"runtime"
)

func main() {
	a := "e"
	b := "é"
	c := "y9999"
	fmt.Printf("%v %v %v\n", a, b, c)

	d := []byte(a) // conversion para byte utilizamos o tipo slice
	e := []byte(b)
	f := []byte(c)
	fmt.Printf("%v %v %v", d, e, f)

	x := 10
	y := 10.0

	fmt.Printf("\n%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)

	// package runtime functions
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

}