package main

import (
	"fmt"
)

var y = 5 // agora a variável y é package-level scope e pode ser lida em qualquer codeblock

func mainVar () {
	z := 30
	qualquercoisa(z)
}

func qualquercoisa (x int) {

	fmt.Println(y) // dentro desse contexto eu não sei oq é a variável y do codeblock main
	fmt.Println(x)
}