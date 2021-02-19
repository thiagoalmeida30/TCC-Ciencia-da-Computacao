package main

import "fmt"

var lista []string
var preco []float32
var produto string
var valor float32
var i int8
var total float32

func main() {
	//loop para criar 2 slices que recebe um imput do usuário e salva no array o nome do produto
	//e outro array que salva o valor
	for i == 0 {
		fmt.Println("Adicionar Produto")
		fmt.Scanln(&produto)
		lista = append(lista, produto)

		fmt.Println("Adcionar Preço")
		fmt.Scanln(&valor)
		preco = append(preco, valor)

		fmt.Println("Novo Produto? ENTER para SIM, 1 para Encerrar Lista")
        fmt.Scanln(&i)
	}
		fmt.Println("Produto", "  Preço")
	//função len retorna o tamanho do slice e joga o valor pra variável indice que será usada como valor do
	//indice do slice para alimentar o loop a seguir
	indice := len(lista) 

	j:=0

	//loop para imprimir um item e seu valor por linha
	for j < indice {
		fmt.Println(lista[j], "   ", preco[j])
		total = total + preco[j] //iteração que usa o indice de cada estado do loop para somar os valores
		j++
	}
	fmt.Printf("Valor Total: %f", total)

}