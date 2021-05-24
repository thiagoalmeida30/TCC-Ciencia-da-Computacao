package main

import "fmt"

func main() {

	var x int = 10
	var y int = x

	x++

	fmt.Println(x)
	fmt.Println(y)
	var ponteiro *int

	ponteiro = &x
	x++

	fmt.Println(ponteiro)
	fmt.Println(x)
	fmt.Println(*ponteiro)

}
