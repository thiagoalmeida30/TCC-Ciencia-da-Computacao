package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
	CPF   string
	sexo  string
}

func main() {

	var novaPessoa pessoa
	fmt.Println("Digite as informações para efetivação do seu cadastr")
	fmt.Println("nome completo")
	fmt.Scanln(&novaPessoa.nome)
	fmt.Println("idade")
	fmt.Sscanln(&novaPessoa.idade)
	fmt.Println("CPF")
	fmt.Sscanln(&novaPessoa.CPF)
	fmt.Println("sexo")
	fmt.Sscanln(&novaPessoa.sexo)

	fmt.Printf("Eu %s, %d anos, sexo %s, portador do CPF %s, autorizo o uso dessas informações pela empresa  cadatro S/A", novaPessoa.nome, novaPessoa.idade, novaPessoa.sexo, novaPessoa.CPF)

}
