package main

import "fmt"

type Persona struct {
	nombre   string
	apellido string
	edad     int
}

func main() {
	persona := Persona{
		nombre:   "Oscar",
		apellido: "Ruiz",
		edad:     24,
	}

	persona.presentar()
}

func (persona Persona) presentar() {
	fmt.Printf("Hola, mi nombre es %v %v\n", persona.nombre, persona.apellido)
	fmt.Printf("Y tengo %v años", persona.edad)
}
