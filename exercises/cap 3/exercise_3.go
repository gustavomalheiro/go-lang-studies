package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%s is %d yo and it is %v!", y, x, z)
	fmt.Println(s)

	fmt.Printf("%T", s)
}