package main

import "fmt"

type vehiculo struct {
	puertas int
	color   string
}

type camion struct {
	vehiculo
	cuatroRuedas bool
}

type sedan struct {
	vehiculo
	lujoso bool
}

func main() {
	camionsito := camion{
		vehiculo: vehiculo{
			puertas: 2,
			color:   "Rojo",
		},
		cuatroRuedas: false,
	}

	sedansito := sedan{
		vehiculo: vehiculo{
			puertas: 4,
			color:   "Blanco",
		},
		lujoso: true,
	}

	fmt.Println(camionsito)
	fmt.Println(camionsito.color)
	fmt.Println(sedansito)
	fmt.Println(sedansito.puertas)
}
