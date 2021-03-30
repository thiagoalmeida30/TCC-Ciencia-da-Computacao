package main

import "fmt"

func inverterSinal(numero *int) {
	*numero = *numero * -1
}

func main() {
	novoNumero := 40

	inverterSinal(&novoNumero)
	fmt.Println(novoNumero)
}