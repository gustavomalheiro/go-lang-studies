package main

import (
	"fmt"
)

var a int
var b float64
var c string
var d bool

func main() {
	fmt.Printf("a: %v, %T\n", a, a) // 0
	fmt.Printf("b: %.1f, %T\n", b, b) // 0.0 
	fmt.Printf("c: %v, %T\n", c, c) // " " -> o nada que tem entre duas aspas
	fmt.Printf("d: %v, %T\n", d, d) // false
}

/*
var x bool // declaração

func main() {
	x = true // inicialização
	
	x = false // atribuição
	fmt.Printf("x: %v, %T", x, x)
	
}
*/