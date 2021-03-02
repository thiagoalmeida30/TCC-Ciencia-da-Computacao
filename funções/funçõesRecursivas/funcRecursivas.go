package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {

	fib := fibonacci(10)
	fmt.Println(fib)

	fib2:=uint(10)
	for i :=uint(1); i <= fib2; i++ {
		fmt.Println(fibonacci(i))
		
		
	}

}
