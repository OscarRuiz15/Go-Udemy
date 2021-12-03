package main

import (
	"encoding/json"
	"fmt"
)

type personita struct {
	Nombre   string `json:"Nombre"`
	Apellido string `json:"Apellido"`
	Edad     int    `json:"Edad"`
}

func main() {
	mar := marshal()
	fmt.Printf("\n%T\n", mar)
	fmt.Println(mar)

	bytes := []byte(mar)
	unmar := unmarshal(bytes)
	fmt.Printf("\n%T\n", bytes)
	fmt.Println(unmar)
}

func unmarshal(mar []byte) []personita {
	var personitas []personita
	err := json.Unmarshal(mar, &personitas)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Toda la data", personitas)

	for i, v := range personitas {
		fmt.Println("\nPERSONA NÚMERO", i)
		fmt.Println(v.Nombre, v.Apellido, v.Edad)
	}

	return personitas
}

func marshal() string {
	p1 := personita{
		Nombre:   "Oscar",
		Apellido: "Ruiz",
		Edad:     24,
	}

	p2 := personita{
		Nombre:   "Alexander",
		Apellido: "Palacio",
		Edad:     25,
	}

	personitas := []personita{p1, p2}
	fmt.Println(personitas)

	bs, err := json.Marshal(personitas)
	if err != nil {
		fmt.Println(err)
	}
	return string(bs)
}