package main

import "fmt"

func main() {
	p1 := struct {
		nombre   string
		apellido string
		edad     int
	}{
		nombre:   "Oscar",
		apellido: "Ruiz",
		edad:     24,
	}
	fmt.Println(p1)
	fmt.Println(p1.nombre, p1.apellido, p1.edad)
}
