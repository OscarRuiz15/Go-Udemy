package main

import "fmt"

func main() {
	n := "Pera"
	switch n {
	case "Manzana":
		fmt.Println("Manzana Roja")
	case "Pera":
		fmt.Println("Pera Verde")
	case "Q":
		fmt.Println("Esta es la q")
	default:
		fmt.Println("Esta es default")
	}
}
}
