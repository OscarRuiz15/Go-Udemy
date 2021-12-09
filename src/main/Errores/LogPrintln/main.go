package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-existe.txt")
	if err != nil {
		log.Println("Ocurrio un error", err)
	}
}
