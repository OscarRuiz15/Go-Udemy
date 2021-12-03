package main

import "fmt"

func main() {
	var x [5]int
	x[3] = 42
	var y [6]int
	fmt.Println(x, len(x))
	fmt.Printf("%T\n", x)
	fmt.Println(y, len(y))
	fmt.Printf("%T", y)
}
