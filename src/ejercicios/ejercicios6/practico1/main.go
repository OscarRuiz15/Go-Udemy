package main

import "fmt"

func main() {
	n := foo()
	x, s := bar()

	fmt.Println(n, x, s)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 1492, "Descubrimiento de América"
}
