package main

import "fmt"

type Persona struct {
	nombre   string
	apellido string
}

type AgenteSecreto struct {
	Persona
	lpm bool
}

func (a AgenteSecreto) presentarse() {
	fmt.Println("Hola, soy", a.nombre, a.apellido)
}

func main() {
	as1 := AgenteSecreto{
		Persona: Persona{
			nombre:   "Oscar",
			apellido: "Ruiz",
		},
		lpm: true,
	}

	as2 := AgenteSecreto{
		Persona: Persona{
			nombre:   "alexander",
			apellido: "palacio",
		},
		lpm: false,
	}

	fmt.Print(as1)
	fmt.Print(as2)

	as1.presentarse()
	as2.presentarse()
}
