package main

import "fmt"

func main() {
	f := retornoFuncion()
	valRetorn := f()
	fmt.Println("El valor que retornó la funcion que fue retornada desde otra funcion es", valRetorn)
}

func retornoFuncion() func() int {
	return func() int {
		fmt.Println("Soy un texto que se imprime desde una funcion que fue retornada desde otra funcion")
		return 1996
	}
}
