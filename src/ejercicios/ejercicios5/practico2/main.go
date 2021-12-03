package main

import "fmt"

type person struct {
	nombre   string
	apellido string
	sabores  []string
}

func main() {
	p1 := person{
		nombre:   "Oscar",
		apellido: "Ruiz",
		sabores:  []string{"Chocolate", "Vainilla", "Fresa"},
	}
	p2 := person{
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

	mapita := map[string]person{
		p1.apellido: p1,
		p2.apellido: p2,
	}
	for key, value := range mapita {
		fmt.Printf("%v ->\n\t%v\n", key, value.nombre)
		for k, v := range value.sabores {
			fmt.Printf("\t\t%v ->\n\t\t%v\n", k, v)
		}
	}
}
