package main

import "fmt"

type vehiculo struct {
	puertas int
	color   string
}

func main() {
	anonimous := struct {
		vehiculo struct {
			puertas int
			color   string
		}
		cuatroRuedas bool
		lujoso       bool
	}{
		vehiculo: vehiculo{
			puertas: 5,
			color:   "Verde",
		},
		cuatroRuedas: false,
		lujoso:       true,
	}

	fmt.Println(anonimous)
	fmt.Println(anonimous.vehiculo)
	fmt.Println(anonimous.cuatroRuedas)
	fmt.Println(anonimous.lujoso)

	s := struct {
		nombre  string
		amigos  map[string]int
		bebidas []string
	}{
		nombre: "Oscar",
		amigos: map[string]int{
			"Carito":  123,
			"Adriana": 456,
			"Condor":  666,
		},
		bebidas: []string{"Agua", "Chicha"},
	}

	fmt.Println(s)
	fmt.Println(s.nombre)
	fmt.Println(s.amigos)
	fmt.Println(s.amigos["Adriana"])
	fmt.Println(s.bebidas)

	for k, v := range s.amigos {
		fmt.Println(k, v)
	}
	for k, v := range s.bebidas {
		fmt.Println(k, v)
	}
}
