package main

import (
	"fmt"
)

func main() {
	for x := 10; x <= 100; x++ {
		resto := x % 4
		fmt.Println(resto)
	}
}