package main

import (
	"fmt"
	"math"
)

func escreverArea(f forma){
	fmt.Printf("A area da forma é %0.2f\n", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

func (u retangulo) area() float64 {
	return u.altura * u.largura
}

type circulo struct {
	raio float64
}

func (u circulo) area() float64 {
	return math.Pi * (u.raio * u.raio)
}

type forma interface {
	area() float64
}

func main() {
r:= retangulo{12.3, 14.2}
fmt.Println(r)
escreverArea(r)

c:= circulo{43}
fmt.Println(c)
escreverArea(c)
}