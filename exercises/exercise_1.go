package main

import (
	"fmt"
)

func main() {
	x := 42
	y := "James Bond"
	z := true

	// unica declaração print
	fmt.Println("Única declaração:")
	fmt.Printf("x: %v\ny: %v\nz: %v\n", x, y, z)

	fmt.Println()

	// multiplas declarações print
	fmt.Println("Múltiplas declarações:")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}

