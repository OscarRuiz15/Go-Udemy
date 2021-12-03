package main

import (
	"fmt"
	"math"
)

type Cuadrado struct {
	longitud float64
}

type Circulo struct {
	radio float64
}

type Forma interface {
	area() float64
}

func main() {
	circ := Circulo{
		radio: 12.345,
	}

	cua := Cuadrado{
		longitud: 15,
	}

	info(circ)
	info(cua)
}

func info(forma Forma) {
	fmt.Println(forma.area())
}

func (cuadrado Cuadrado) area() float64 {
	fmt.Println("Area de un cuadrado")
	return cuadrado.longitud * cuadrado.longitud
}

func (circulo Circulo) area() float64 {
	fmt.Println("Area de un circulo")
	return math.Pi * circulo.radio * circulo.radio
}