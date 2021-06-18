package main

import (
	"fmt"
)

func main() {
	
	/*
	x := 0 

	for x <= 10 {
		fmt.Println(x)
		x++
	}
	*/
	teste()
}

func teste() {
	x := 0

	for {
		if x < 10 {
			x++
			fmt.Println(x)
		} else {
			break
		}
	}
	
}