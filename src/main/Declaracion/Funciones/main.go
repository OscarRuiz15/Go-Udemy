package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("Mi primera expresion funcion")
	}

	f()

	g := func(x int) {
		fmt.Println("El año en que se descubrio America fue", x)
	}

	g(1492)
}
