package main

import (
	"fmt"
)

func main() {
	agenda := [][]string {
			 {
				"gustavo",
				"malheiro",
				"jogar",
			},
			{
				"juliana",
				"garcia",
				"desenhar",
			},
			{
				"juan",
				"latorre",
				"ouvir kpop",
			},
	}

	for _, v := range agenda {
		fmt.Println(v)
	}

	fmt.Println("")
}