package main

import "fmt"

func somar(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero

	}
	return total
}

func main() {
	soma := somar(2, 4, 35, 54, 12, 1, 32, 3)

	fmt.Println(soma)
}