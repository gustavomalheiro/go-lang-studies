package main

import (
	"fmt"
)

func mainGopher () {

	// tipagens automáticas que dependem do valor que foi atribuido
	// gopher
	x := 10
	y := "olá bom dia"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)

	// operador de atribuição
	x = 20 
	fmt.Printf("x: %v, %T", x, x)

}