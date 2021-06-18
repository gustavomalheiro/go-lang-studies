package main

import (
	"fmt"
)

const (
	x = 10 // esse ficará indefinido até ele ser usado
) 

var (
	y = 10 // no momento que o compilador ler essa linha, ele vai tipar como int
	c int
	d float64
)

func main() {
	
	d = x
	
	fmt.Printf("%v", d)
}