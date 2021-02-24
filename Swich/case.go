package main

import "fmt"

func diaDaSemana(dia int) string {
	switch dia {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "quarta-feira"
	case 5:
		return "quinta-feira"
	case 6:
		return "sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

func main() {
	data := diaDaSemana(3)
	fmt.Println(data)
}
