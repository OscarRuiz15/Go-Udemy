package main

import "fmt"

func main() {
	fmt.Println("Funcion anonima se ejecutó")

	func() {
		fmt.Println("La función se ejecutó")
	}()

	func(x int) {
		fmt.Println("El significado de la vida es:", x)
	}(5)
}
