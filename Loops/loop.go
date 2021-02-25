package main

import (
	"fmt"
	"time"
)



func main() {
	i := 0

	for i <= 10 {
		fmt.Println(i)
		i++
		time.Sleep(time.Second)
	}
	fmt.Println("fim do loop")

	nome := [4]string {"Pedro", "Thiago", "joão", "No barquinho"} 
	for _, valor := range nome{
		fmt.Println(valor)
		time.Sleep(time.Second)

	}

	
}
