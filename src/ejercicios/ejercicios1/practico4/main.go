package main

import "fmt"

type numero int

func main() {
	var x numero
	fmt.Println(x)
	fmt.Printf("x: %T\n", x)
	x = 42
	fmt.Println(x)
}
