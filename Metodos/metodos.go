package main

import "fmt"

type usuario struct {
    nome string
    idade int
}
func (u usuario) salvar(){
    fmt.Printf("salvando o usuário %s no banco de dados \n", u.nome)
}
func (u usuario) maiorDeIdade(){
    if u.idade >= 18 {
        maioridade := true
    fmt.Println(maioridade)
    fmt.Printf("%s é maior de idade \n", u.nome)
    }else {
        maioridade := false
    fmt.Println(maioridade)
    fmt.Printf("%s é menor de idade \n", u.idade)
    }
    
    }
    
    func (u *usuario) fazerAniversario(){
        u.idade++
        fmt.Println(u.idade)
    }

func main() {

usuario1 := usuario{"thiago", 32}
usuario1.salvar()
usuario1.maiorDeIdade()
usuario1.fazerAniversario()
}
