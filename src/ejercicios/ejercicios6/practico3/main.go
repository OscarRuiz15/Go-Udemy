package main

import "fmt"

func main() {
	defer conDefer()
	sinDefer()
}

func conDefer() {
	defer func() {
		fmt.Println("Una funcion diferida dentro de otra, esta deberia estar de ultima")
	}()
	fmt.Println("Funcion con defer de primera, ejecuta de ultima ?")
}

func sinDefer() {
	fmt.Println("Funcion sin defer de segunda, deberia ejecutar de primera")
}
