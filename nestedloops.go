package main

import (
	"fmt"
)

func main() {
	for horas := 0; horas <= 12; horas++ {
		fmt.Println("Hora:", horas)
		for x := 0; x <= 60; x++ {
			fmt.Printf("%d ", x)
		}
		fmt.Println("")
	}	
}