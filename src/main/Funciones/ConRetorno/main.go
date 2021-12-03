package main

import "fmt"

func main() {
	fmt.Println("\nFunciones con retorno")
	value := returnValue(5)
	fmt.Println(value)

	value1, value2 := doubleReturn(7)
	fmt.Println(value1, value2)

	value3, _ := doubleReturn(3)
	fmt.Println(value3)
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func returnValue(a int) int {
	return a * 2
}
