package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("asdasd")
	if err != nil {
		log.Panicln(err)
	}
}
