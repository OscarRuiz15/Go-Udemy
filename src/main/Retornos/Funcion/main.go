package main

import "fmt"

func main() {
	x := funcionRetornaFuncion()
	fmt.Printf("%T\n", x)

	i := x()
	fmt.Println(i)

	fmt.Println(funcionRetornaFuncion()())
}

func funcionRetornaFuncion() func() int {
	return func() int {
		return 1492
	}
}