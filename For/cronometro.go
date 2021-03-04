package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	var j int16
	fmt.Println("INFORME O TEMPO EM SEGUNDOS PARA INICIAR O CRONOMETRO")
	fmt.Scanln(&j)
	i := j
	fmt.Println("Contagem regressiva...")
	time.Sleep(time.Second)
	for i > 0 {

		fmt.Println(i)
		i--
		time.Sleep(time.Second)

	}

	fmt.Println("Tempo esgotado")

	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
