package main

import "fmt"

func recuperarExecucao()  {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
		
	}
	
		fmt.Println("Está em recuperação")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A MEDIA É EXATAMENTE 6")
}

func main() {

	resultado := alunoEstaAprovado(6, 6)
	fmt.Println(resultado)
}
