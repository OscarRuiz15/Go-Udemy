package main

import "fmt"

func main() {
	//unbuffered channel (canal sin buffer)
	//El tipo de valor que el canal va a estar enviando y recibiendo (int)
	ca := make(chan int, 2)

	ca <- 42
	ca <- 43

	fmt.Println(<-ca)
	fmt.Println(<-ca)

}
