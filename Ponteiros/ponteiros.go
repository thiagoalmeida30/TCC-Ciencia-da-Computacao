package main

import "fmt"

func main() {

	var variavel1 int = 10
	var variavel2 int = variavel1

	variavel1++

	fmt.Println(variavel2)
	fmt.Println(variavel1)

	//PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA

	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3)
	fmt.Println(ponteiro)

	fmt.Println(*ponteiro)
}
