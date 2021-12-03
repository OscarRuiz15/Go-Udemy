package main

import "fmt"

// El tipo agentSecret recibe el metodo introduce, y como el metodo introduce es uno de los metodos
// en la interfaz human, entonces agentSecret tambien es de tipo human, es decir, implementa la
// interfaz human

type human interface {
	introduce()
}

type person struct {
	firstName string
	lastName  string
}

type agentSecret struct {
	person
	lpm bool
}

type manzana int

func (p person) introduce() {
	fmt.Println("Hola (Persona), mi nombre es ", p.firstName, p.lastName)
}

func (as agentSecret) introduce() {
	fmt.Println("Hola (Agente secreto), mi nombre es ", as.firstName, as.lastName)
}

func bart(h human) {
	fmt.Println("Fui pasado a la funcion bart", h)
	switch h.(type) {
	case person:
		fmt.Println("Ah, soy una persona", h.(person).firstName)
	case agentSecret:
		fmt.Println("Ah, soy un agente secreto", h.(agentSecret).firstName)
	}
}

func main() {
	// Conversion
	var x manzana = 42
	fmt.Println("X", x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println("Y", y)
	fmt.Printf("%T\n", y)

	as1 := agentSecret{
		person: person{
			firstName: "Oscar",
			lastName:  "Ruiz",
		},
		lpm: true,
	}

	p1 := person{
		firstName: "Alexander",
		lastName:  "Palacio",
	}

	as1.introduce()
	p1.introduce()

	bart(as1)
	bart(p1)
}
