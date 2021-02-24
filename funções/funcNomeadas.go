package main

import "fmt"

func cauculosMatematicos(n1, n2 int)(soma, subtracao int) {
    soma = n1 + n2
    subtracao = n1 - n2
    return
}

func main() {
soma1, subtracao1 := cauculosMatematicos(30, 20)

fmt.Printf("A soma é: %d \n", soma1)
fmt.Printf("A subtração é: %d \n", subtracao1)
}

/*Funções com retorno nomeado, são funções onde o retorno já declara 
a variável no parâmetro aí invés de apenas declarar o tipo do retorno*/
