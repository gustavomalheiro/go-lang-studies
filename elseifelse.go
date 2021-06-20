package main

import (
	"fmt"
)

func main() {
	
	if x := 91; x > 100 {
		fmt.Println("Hello!")
	} else if x < 100 && x > 90 {
		fmt.Println("Goodbye")
	} else {
		fmt.Println("OSHIETEOSHIETE")
	}

}