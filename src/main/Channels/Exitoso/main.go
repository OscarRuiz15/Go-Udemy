package main

import "fmt"

func main() {
	//unbuffered channel (canal sin buffer)
	//El tipo de valor que el canal va a estar enviando y recibiendo (int)
	ca := make(chan int)

	go func() {
		ca <- 42
	}()

	fmt.Println(<-ca)

}
