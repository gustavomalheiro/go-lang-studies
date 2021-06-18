package main

import (
	"fmt"
)

func main() {
	x := "oi"
	y := "bom dia"
	fmt.Println(x) // repare que há uma linha em branco no final
	fmt.Println(y)

	//Sprintln
	z := fmt.Sprintln(x, y) // atualização: coloca espaço entre os operandos e uma nova linha é adicionada
	fmt.Println(z)
}

/*
func main() {
	// interpreted string literal
	x := "oi bom dia\ncomo vai?\tespero \"que\" tudo bem."
	fmt.Println(x)
	// essa string x é interpretada. a função passa pelos bytes e vai vendo oq temos, como os \n, \t ou \", que tem interpretações próprias.

	// nós temos também os raw string literals. a diferença entre um e outro, é que um eu uso " e outro uso `
	y := `"oi bom dia\ncomo vai?                     \tespero \"que\"                       tudo bem."`
	fmt.Println(y)
	// ele vai fazer basicamente que nada aqui dentro seja interpretado, e tudo vai ficar CRU, da maneira que eu digitei.
}
*/