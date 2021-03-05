package main

import "fmt"

func main() {

	func() {
		defer fmt.Println("teste defer 1")
		fmt.Println("teste defer 2")
	}()
		defer fmt.Println("A média foi calculada o resultado será retornado")
		fmt.Println("Entrando na função media")
	resultado:=func(n1, n2 float32) bool {

		media:= (n1 + n2)/2

		if media >= 5 {
			return true
		}
			return false
	
	}(7,8)

	fmt.Println(resultado)

}
