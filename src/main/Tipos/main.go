package main

import "fmt"

type dinero int

var osc dinero

func main() {
	fmt.Println("\nCreacion de tipos")
	osc = 1000
	fmt.Println(osc)
	fmt.Printf("%T", osc)
}
