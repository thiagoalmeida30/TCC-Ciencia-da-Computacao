package main

import "fmt"

func main() {
	fmt.Println("ponto de partida")

	aplicacao := app.Gerar()
	aplicacao.Run(os.Args)
}