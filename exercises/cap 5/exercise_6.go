package main

import (
	"fmt"
)

/*
minha solução

const (
	x = iota + 2022
	y
	z
	a
)
*/

/*
solução da professora que achei melhor, pois abrange mais fundamentos da linguagem
*/

const (
	_ = 2021 + iota // não usaremos o ano atual, por isso descartamos com o _
	x
	y
	z
	a
)

func main() {
	fmt.Println(x, y, z, a)
}