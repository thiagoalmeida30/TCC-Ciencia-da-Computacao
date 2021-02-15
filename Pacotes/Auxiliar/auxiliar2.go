package auxiliar

import "fmt"

//imprime um texto na tela
//função escrever de forma miuscula representa uma função tal como private
//GO não possui declaração explicita de privado ou publico sendo definido
//pela primeira letra da função, Maiúscula representa uma função pública
// e a minúscula representa a função privada assim como no auxiliar2.go
func escrever2() {
	fmt.Println("escrevendo do pacote auxiliar 2")
}