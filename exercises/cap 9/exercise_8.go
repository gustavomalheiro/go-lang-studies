package main

import (
	"fmt"
)

func main() {
	pessoa := map[string][]string {
		"malheiro_gustavo" : {"Jogar", "Ouvir Música"},
		"garcia_juliana" : {"Jogar", "Desenhar"},
		"latorre_juan" : {"Ouvir KPOP", "Desenhar"},
		
	}

	for key, value := range pessoa {
		fmt.Println(key)
		for i, h := range value {
			fmt.Println("\t", i, " - ", h)
		}
	}
}