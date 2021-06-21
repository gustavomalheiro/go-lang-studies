package main

import (
	"fmt"
)

func main() {

	pessoa := map[string][]string{
		"malheiro_gustavo": {
			"Jogar", 
			"Ouvir Música",
		},
		"garcia_juliana": {
			"Desenhar",
			"Jogar",
			"Ler",
		},
		"latorre_juan": {
			"Desenhar",
			"Ouvir KPOP",
		},
	}

	pessoa["garcia_miguel"] = []string{"Jogar Lol", "Jogar Dark Souls 3"}

	for key, value := range pessoa {
		fmt.Println(key)
		for i, v := range value {
			fmt.Println("\t", i, v)
		}
	}


}