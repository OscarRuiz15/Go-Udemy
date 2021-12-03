package main

import "fmt"

func main() {
	n1 := factorial(5)
	fmt.Println(n1)

	n2 := cicloFac(6)
	fmt.Println(n2)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func cicloFac(n int) int {
	fact := 1
	for ; n > 0; n-- {
		fact *= n
	}

	return fact
}
