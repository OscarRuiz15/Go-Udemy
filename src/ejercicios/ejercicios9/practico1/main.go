package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Comienza la ejecución")

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		fmt.Println("Primera go rutine")
		wg.Done()
	}()

	go func() {
		fmt.Println("Segunda go rutine")
		wg.Done()
	}()

	fmt.Println("A punto de terminar")

	wg.Wait()
}
