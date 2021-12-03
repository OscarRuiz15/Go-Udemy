package main

import "fmt"

func main() {
	fmt.Println("\nOperadores Logicos")
	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma", result)

	// Resta
	result = y - x
	fmt.Println("Resta", result)

	// Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion", result)

	// Division
	result = y / x
	fmt.Println("Division", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo", result)

	// Incremental
	x++
	fmt.Println("Incrementa x", x)

	// Decremental
	x--
	fmt.Println("Decremental x", x)
}
