package main

import (
	"encoding/json"
	"fmt"
)

type personita struct {
	Nombre   string `json:"Nombre"`
	Apellido string `json:"Apellido"`
	Edad     int    `json:"Edad"`
	Dichos   []string
}

func main() {
	s := `[{"Nombre":"Eduar","Apellido":"Tua","Edad":32,"Dichos":["Cachicamo diciéndole a morrocoy conchudo","La mona, aunque se vista de seda, mona se queda","Poco a poco se anda lejos"]},{"Nombre":"Carlos","Apellido":"Pérez","Edad":27,"Dichos":["Camarón que se duerme se lo lleva la corriente","A ponerse las alpargatas que lo que viene es joropo","No gastes pólvora en zamuro"]},{"Nombre":"M","Apellido":"Hmmmm","Edad":54,"Dichos":["Ni lava ni presta la batea","Hijo de gato, caza ratón","Más vale pájaro en mano que cien volando"]}]`
	fmt.Println(s)

	bytes := []byte(s)
	var personitas []personita
	err := json.Unmarshal(bytes, &personitas)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Toda la data", personitas)

	for i, persona := range personitas {
		fmt.Println("Persona #", i+1)
		fmt.Println("\t", persona.Nombre, persona.Apellido, persona.Edad)
		for _, dichos := range persona.Dichos {
			fmt.Println("\t\t", dichos)
		}
	}
}
