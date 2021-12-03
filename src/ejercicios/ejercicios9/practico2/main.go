package main

import "fmt"

type Persona struct {
	Nombre   string
	Apellido string
	Edad     int64
}

type Humano interface {
	hablar()
}

func (persona *Persona) hablar() {
	dialogo := fmt.Sprintf("%s %s %s %s %v %s", "Hola, mi nombre es", persona.Nombre, persona.Apellido, "Y tengo", persona.Edad, "años")
	fmt.Println(dialogo)
}

func diAlgo(humano Humano) {
	humano.hablar()
}

func main() {
	persona := Persona{
		"Oscar",
		"Ruiz",
		25,
	}

	diAlgo(&persona)
}
