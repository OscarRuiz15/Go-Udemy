package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int)
	cw := make(chan<- int)
	fmt.Println("-------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("c\t%T\n", cr)
	fmt.Printf("c\t%T\n", cw)

	cr = c
	cw = c
	fmt.Println("-------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("c\t%T\n", cr)
	fmt.Printf("c\t%T\n", cw)

	// convertir de general a especifico
	fmt.Println("-------")
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Printf("c\t%T\n", (chan<- int)(c))

	// especifico a general (no convierte)
	//fmt.Println("-------")
	//fmt.Printf("c\t%T\n", (chan int)(cr))
	//fmt.Printf("c\t%T\n", (chan int)(cw))
}
