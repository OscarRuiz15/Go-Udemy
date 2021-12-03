package main

import "fmt"

func main() {
	fmt.Println("\nDeclaracion funciones")
	normalFunction("Hola mundo")
	normalFunction("Hola mundo 2")
	normalFunction("Hola mundo 3")
	tripleArgument(1, 2, "Hola")
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func normalFunction(message string) {
	fmt.Println(message)
}