package main

import "fmt"

func main() {
	mapita := map[string][]string{
		"eduar_tua":    {"computadoras", "montaña", "playa"},
		"carlos_ramon": {"leer", "comprar", "musica"},
		"juan_bimba":   {`helado`, `pintar`, `bailar`},
	}

	mapita["oscar_ruiz"] = []string{"Uno", "Dos", "Tres"}
	delete(mapita, "carlos_ramon")

	for key, value := range mapita {
		fmt.Println(key)
		for key2, value2 := range value {
			fmt.Println("\t", key2, value2)
		}
	}
}
