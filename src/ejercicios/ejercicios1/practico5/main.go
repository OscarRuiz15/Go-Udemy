package main

import "fmt"

type num int

var equis num
var ye int

func main() {
	fmt.Println(equis)
	fmt.Printf("equis: %T\n", equis)
	equis = 42
	fmt.Println(equis)

	ye = int(equis)
	fmt.Println(ye)
	fmt.Printf("ye: %T\n", ye)
}
