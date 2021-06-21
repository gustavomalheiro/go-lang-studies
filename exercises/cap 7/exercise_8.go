package main

import (
	"fmt"
)

func main() {
	
	x := 21

	switch {
		case x > 20:
			fmt.Println("x é maior que 20!")
		case x < 20 && x > 10:
			fmt.Println("x é menor que 20!")
		case x == 10:
			fmt.Println("x é 10!")
		default:
			fmt.Println("x é x")
	}
}