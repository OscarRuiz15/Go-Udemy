package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
}

func main() {
	person := Persona{
		nombre: "Oscar",
		edad:   24,
	}

	fmt.Println(person)
	cambiame(&person)
	fmt.Println(person)
}

func cambiame(persona *Persona) {
	persona.nombre = "Alexander"
	persona.edad = 25
}
