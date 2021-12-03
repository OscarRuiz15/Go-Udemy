package main

import "fmt"

func main() {
	func(x int) {
		operacion := x * x
		fmt.Printf("%v elevado al cuadrado es %v", x, operacion)
	}(5)
}
