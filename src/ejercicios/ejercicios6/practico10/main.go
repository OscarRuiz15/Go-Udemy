package main

import "fmt"

func main() {
	fmt.Println("re")
	re := funcclosure()
	fmt.Println(re())
	fmt.Println(re())
	fmt.Println(re())
	fmt.Println(re())

	fmt.Println("ra")
	ra := funcclosure()
	fmt.Println(ra())
	fmt.Println(ra())
	fmt.Println(ra())
	fmt.Println(ra())
	fmt.Println(ra())
	fmt.Println(ra())
	fmt.Println(ra())
	fmt.Println(ra())
}

func funcclosure() func() int {
	var x = 0
	return func() int {
		x++
		return x
	}
}
