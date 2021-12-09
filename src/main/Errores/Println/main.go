package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("no-existe.txt")
	if err != nil {
		fmt.Println("Ocurrio un error", err)
	}
}
