package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	matricula string
	curso     string
	turno     string
}

func main() {

	estudante1 := estudante{pessoa{"thiago", "almeida", 32, 190}, "2014101648", "Ciencia da Computação", "noite"}
	fmt.Println(estudante1)

	fmt.Println(estudante1.nome)

	pessoa1 := pessoa{"Caio", "braga", 29, 178}

	estudante2 := estudante{pessoa1, "2016141311", "Design", "manhã"}

	fmt.Println(estudante2)

}
