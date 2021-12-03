package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	sabores  []string
}

func main() {
	p1 := persona{
		nombre:   "Oscar",
		apellido: "Ruiz",
		sabores:  []string{"Chocolate", "Vainilla", "Fresa"},
	}
	p2 := persona{
		nombre:   "Alexander",
		apellido: "Palacio",
		sabores:  []string{"Maracuya", "Mora", "Ron con pasas"},
	}

	fmt.Println(p1)
	for key, value := range p1.sabores {
		fmt.Println("\t", key, value)
	}

	fmt.Println(p2)
	for key, value := range p2.sabores {
		fmt.Println("\t", key, value)
	}
}
