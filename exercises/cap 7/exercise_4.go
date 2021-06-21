package main

import "fmt"

func main() {

	ano := 2000
	anofinal := 2021

	for {
		if (ano > anofinal) {
			break	
		} 
		fmt.Println(ano)
		ano++
	}
}