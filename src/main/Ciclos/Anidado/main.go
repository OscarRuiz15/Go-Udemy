package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("i", i)
		for a := 0; a < 3; a++ {
			fmt.Println("\t\ta", a)
		}
		fmt.Println()
	}
}
