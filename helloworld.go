package main

import (
	"fmt"
)

func main() {

	x := 16 //numeros
	y := "strings" // letras (cadeia de)
	z := true // v ou f

	fmt.Println(x, y, z)

	_, errorTest := fmt.Println("Hello world!")
	fmt.Println(errorTest) 
}