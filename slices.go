package main

import (
	"fmt"
)

func main() {
	array := [5]int{1,2,3,4,5}
	fmt.Println(array)
	slice := []int{1,2,3,4,5}
	fmt.Println(slice)

	fmt.Printf("%T\n", slice)

	// adicionando um elemento ao final do slice
	slice = append(slice, 6)
	fmt.Println(slice)

	// adicionando um novo valor a uma posição do slice
	fmt.Println(slice[3])
	slice[3] = 348756
	fmt.Println(slice[3])

	fmt.Println("")

	fmt.Println("")

	arraySubjacente()
}

func forange() {
	slice := []string{"banana", "maçã", "jaca", "pêssego"}
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	for indice, valor := range slice {
		fmt.Println("No índice", indice, "temos o valor:", valor)
	}

	slice = append(slice, "melancia")

	fmt.Println("")

	for indice, valor := range slice {
		fmt.Println("No índice", indice, "temos o valor:", valor)
	}

	for indice, _ := range slice {
		fmt.Println("índices desse slice: ", indice)
	}

	for _, valor := range slice {
		fmt.Println("Um dos valores desse slice é: ", valor)
	}

	total := 0

	for _, valor := range numeros {
		total += valor
	}

	println("O valor total é:", total)
}

func fatiandofatias() {

	sabores := []string{"chocolate", "morango", "leite ninho", "formigueiro", "aipim"}

	// 3 até o final
	fatia := sabores[3:] 

	//0 até 0 3
	fatia2 := sabores[:3]

	// acessando os valores do vetor sem range
	for x := 0; x < len(sabores); x++ {
		fmt.Println("Sabores: ", sabores[x])
	}

	// deletando valores do vetor
	sabores = append(sabores[:2], sabores[3:]...) // tudo que vem antes, até ele e tudo que vem depois
	fmt.Println(sabores)

	fmt.Println(fatia, fatia2)
}

func anexar() {
	umaslice := []int{1, 2, 3, 4}
	outraslice := []int{9, 10, 11, 12}

	umaslice = append(umaslice, 5, 6, 7, 8 )

	fmt.Println(umaslice)

	umaslice = append(umaslice, outraslice...)

	fmt.Println(umaslice)
}

func funcaoMake() {
	slice := make([]int, 5, 10)
	
	// ocupando o comprimento definido do slice
	slice[0], slice[1], slice[2], slice[3], slice[4] = 1, 2, 3, 4, 5

	// agora, adicionando um valor além do comprimento (só possível com a função append)
	slice = append(slice, 6)
	slice = append(slice, 7) 
	slice = append(slice, 8) 
	slice = append(slice, 9) 
	slice = append(slice, 10) 

	fmt.Printf("%v \n%v %v ", slice, len(slice), cap(slice))

	// indo acima da minha capacidade, a capacidade vai ser dobrada para acomodar o meu número
	slice = append(slice, 11) 

	fmt.Printf("\n%v \n%v %v ", slice, len(slice), cap(slice))
}

func fatiaDeFatias() {
	ss := [][]int{
			[]int{1, 2, 3, 4, 5, 6}, // 0
			[]int{7, 8, 9, 10, 11, 12}, //1
			[]int{13, 14, 15, 16, 17, 18}, //2
	}

	fmt.Println(ss[1][1]) // 8
	fmt.Println(ss[2][0]) // 13   

}

func arraySubjacente() {
	primeiroslice := []int{1, 2, 3, 4, 5}
	segundoslice := append(primeiroslice[:2], primeiroslice[4:]...)

	fmt.Println(segundoslice)

}