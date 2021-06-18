package main

import (
	"fmt"
)

const x string = "ohayou!"
const y = "ohayou!"

func main() {
	fmt.Printf("%s, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)
}