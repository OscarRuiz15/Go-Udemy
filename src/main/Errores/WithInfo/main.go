package main

import (
	"errors"
	"fmt"
	"log"
)

var errMath = errors.New("De matematica elemental: No hay raiz cuadrada real de un numero negativo")

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		//return 0, errors.New("De matematica elemental: No hay raiz cuadrada real de un numero negativo")
		//return 0, errMath
		return 0, fmt.Errorf("De matematica elemental: No hay raiz cuadrada real de un numero negativo: %v", f)
	}
	return 42, nil
}
