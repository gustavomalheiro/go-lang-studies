package main

import (
	"fmt"
)

func main() {

	esporteFavorito := "trilha de node"

	switch esporteFavorito {
		case "trilha de react":
			fmt.Println("inscrito na trilha de react")
			fallthrough
		case "trilha de node":
			fmt.Println("inscrito na trilha de node")
		case "trilha de elixir":
			fmt.Println("inscrito na trilha de elixir")
		case "trilha de flutter":
			fmt.Println("incrito na trilha de flutter")
		default:
			fmt.Println("inscrito na trilha de react native")
	}
}