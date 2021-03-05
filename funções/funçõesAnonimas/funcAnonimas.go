package main

import "fmt"

func main() {

	func() {
		fmt.Println("Função anônima")
	}()

	func(texto string){
		fmt.Println(texto)

	}("Função Anônima2")

	retorno:=func(texto2 string)string{
		return fmt.Sprintf("Recebido --> %s", texto2)
	}("Função Anônima 3")

	fmt.Println(retorno)
}
