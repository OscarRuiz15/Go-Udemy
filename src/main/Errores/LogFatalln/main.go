package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("sin-archjivo.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

func foo() {
	fmt.Println("Cuando os.Exit() es llamada, las funciones diferidas no corren")
}

/*
Las funciones de Fatal llaman a os.Exit(1) después de escribir el mensaje del log
*/

// Fatalln es equivalente a Println() seguido por una llamada a os.Exit(1)
