package main

import "fmt"

//basicamente a função init será uma função que sempre será executada antes da função main, pode ser usada pra iniciar uma variavel global por exemplo, ou um setup
var n int

func init(){
    fmt.Println("função init")
    n =10
    fmt.Println(10)
}
func main() {

fmt.Println("função main")

}
