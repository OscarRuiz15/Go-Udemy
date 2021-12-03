package main

import "fmt"

func main() {
	fun := func(x ...int) int {
		lon := len(x)
		if lon == 0 {
			return 0
		} else if lon == 1 {
			return x[0]
		} else {
			return x[0] + x[lon-1]
		}
	}

	x := []int{1, 2, 3, 4, 5}
	retorno := recibeFuncion(fun, x...)
	fmt.Println("El retorno de todo eso es", retorno)
}

func recibeFuncion(f func(x ...int) int, x ...int) int {
	dev := f(x...) + 1
	return dev
}
