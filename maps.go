package main

import (
	"fmt"
)

func main() {

	amigos := map[string]int{
		"alfredo": 551234,
		"joana": 9996674,
	}

	fmt.Println(amigos) // mostra tudo que tem no map
	fmt.Println(amigos["joana"]) // mostra apenas o numero que está acoplado a chave joana
	
	// colocando valores novos no map
	amigos["gopher"] = 444444

	fmt.Println(amigos)
	fmt.Println(amigos["gopher"])

	// e se eu quiser o telefone do romário? (não está no map)
	fmt.Println(amigos["romário"]) // 0

	// comma ok idiom
	sera, ok := amigos["fantasma"]

	fmt.Println(sera, ok) 

	if !ok {
		fmt.Println("não tem!")
	} else {
		fmt.Println(sera)
	}

	rangeEmMapas()

	

}

func rangeEmMapas() {
	
	fmt.Println("")

	qualquercoisa := map[int]string{
		123: "muito legal",
		98: "menos legal um pouquinho",
		983: "esse é massa",
		18: "idade de ir pra festa",
	}

	fmt.Println(qualquercoisa)

	for key, value := range qualquercoisa {
		fmt.Println(key, value)
	}

	// deletando elementos de um map
	delete(qualquercoisa, 123)

	fmt.Println(qualquercoisa)

	
}