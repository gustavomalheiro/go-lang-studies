package main

import (
	"fmt"
)

func main() {
	numeros := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fatia1 := numeros[:3]

	fmt.Println(fatia1)

	fatia2 := numeros[4:]

	fmt.Println(fatia2)

	fatia3 := numeros[1:7]

	fmt.Println(fatia3)

	fatia4 := numeros[2:9]

	fmt.Println(fatia4)

	//desafio:
	fatia5 := numeros[2:len(numeros)-1]

	fmt.Println(fatia5)

	
}