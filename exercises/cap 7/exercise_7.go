package main

import (
	"fmt"
)

func main() {
	querocafe := false
	queroagua := false

	if querocafe {
		fmt.Println("vai na cozinha pegar")
	} else if !(querocafe) && queroagua {
		fmt.Println("então toma água")
	} else {
		fmt.Println("poucas ideias")
	}
}