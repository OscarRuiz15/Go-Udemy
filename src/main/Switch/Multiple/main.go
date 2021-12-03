package main

import "fmt"

func main() {
	n := "Pera"
	switch n {
	case "Manzana", "Pera", "Mango":
		fmt.Println("Varias frutas")
	case "M":
		fmt.Println("m")
	case "Q":
		fmt.Println("Esta es la q")
	default:
		fmt.Println("Esta es default")
	}
}
