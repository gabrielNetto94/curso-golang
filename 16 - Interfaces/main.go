package main

import (
	"fmt"
	"math"
)

type Retangulo struct {
	altura  float64
	largura float64
}

func (r Retangulo) area() float64 {
	return r.altura * r.largura
}

type Circulo struct {
	raio float64
}

func (c Circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

type forma interface {
	area() float64
}

func printArea(f forma) {
	fmt.Printf("A área da forma é %0.2f \n", f.area())
}

func main() {
	r := Retangulo{10, 12}
	printArea(r)

	c := Circulo{10}
	printArea(c)
}
