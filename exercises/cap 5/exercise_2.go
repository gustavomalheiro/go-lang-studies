package main

import (
	"fmt"
)

const (
	_ = iota * 2
	x
	y
)

func main() {

	a := (x == y)
	b := (x != y)
	c := (x <= y)
	d := (x < y)
	e := (x >= y)
	f := (x > y)

	fmt.Println(a, b, c, d, e, f)
}