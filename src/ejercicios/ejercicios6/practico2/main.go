package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum1 := funcionA(x...)
	sum2 := funcionB(x)
	fmt.Println("Suma 1 ", sum1)
	fmt.Println("Suma 2 ", sum2)
}

func funcionA(x ...int) int {
	aux := 0
	for i := 0; i < len(x); i++ {
		aux += x[i]
	}

	return aux
}

func funcionB(x []int) int {
	aux := 0
	for _, v := range x {
		aux += v
	}

	return aux
}
