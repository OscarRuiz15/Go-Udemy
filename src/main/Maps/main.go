package main

import "fmt"

func main() {
	m := map[string]int{
		"Batman": 32,
		"Robin":  27,
	}

	fmt.Println(m)
	fmt.Println(m["Batman"])
	fmt.Println(m["Batmann"]) //Si se pasa una llave que no existe, devuelve 0
	if v, ok := m["Batman"]; ok {
		fmt.Println("Existe, la edad es", v)
	} else {
		fmt.Println("No existe")
	}

	// Agregar elementos
	fmt.Println()
	m["Oscar"] = 24
	for key, value := range m {
		fmt.Println(key, value)
	}

	// Eliminar elementos
	fmt.Println()
	fmt.Println(m)
	delete(m, "Robin")
	fmt.Println(m)
	if v, ok := m["Batman"]; ok {
		fmt.Println("Se borró la llave con valor", v)
		delete(m, "Batman")
	} else {
		fmt.Println("No existe")
	}
	fmt.Println(m)
}
