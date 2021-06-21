package main

import (
	"fmt"
)

func main() {
	pessoa := map[string][]string {
		"malheiro_gustavo": {
			"jogar",
			"testar",
		},
		"garcia_juliana": {
			"ouvir música",
			"jogar minecraft",
		},
		"latorre_juan": {
			"ouvir kpop",
			"estudar segurança da informação",
		},
	}

	pessoa["garcia_miguel"] = []string{"jogar lol", "jogar minecraft"}

	for key, value := range pessoa {
		fmt.Println(key)
		for i, v := range value {
			fmt.Println("\t", i, v)
		}
	}

	delete(pessoa, "garcia_miguel")

	fmt.Println("")

	for key, value := range pessoa {
		fmt.Println(key)
		for i, v := range value {
			fmt.Println("\t", i, v)
		}
	}
}