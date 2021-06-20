package main

import (
	"fmt"
)

func main() {

	x := 10

	if x == 10 {
		fmt.Println("Olá!")
	}

// posso usar negação com o operador não

	if !(x == 10) {
		fmt.Println("Olá")
	}

	declaracaoNaInicializacao()
}

func declaracaoNaInicializacao() {
	if x := 10; x == 10 {
		fmt.Println("Olá")
	}
	// não é muito usado, mas faz parte da linguagem
}