package main

import (
	"fmt"
)


func main() {

	usuario := map[string]string{
		"nome":      "Thiago",
		"sobrenome": "Almeida",
	}
	fmt.Println(usuario["nome"])

	fmt.Println(usuario)

	delete(usuario, "nome")

	fmt.Println(usuario)


}