package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for key, value := range arr {
		fmt.Println(key, value)
	}
	fmt.Printf("%T", arr)
}
