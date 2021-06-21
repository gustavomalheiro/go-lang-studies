package main

import (
	"fmt"
)

func main() {
	 fatia := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	 
	 for i, v := range fatia {
		 fmt.Println(i, v)
	 }

	 fmt.Printf("%T", fatia)
}