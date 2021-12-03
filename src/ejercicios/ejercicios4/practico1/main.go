package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	for key, value := range arr {
		fmt.Println(key, value)
	}
	fmt.Printf("%T", arr)
}
