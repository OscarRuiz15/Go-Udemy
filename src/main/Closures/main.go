package main

import "fmt"

func main() {
	// Es una funcion anonima que se retorna, con esa funcion se retorna una copia del entorno,
	// por ende si se hacen varios llamados usando el mismo identificador, se hace apunte al mismo
	// lugar en memoria, entonces como se mantiene una copia del entorno, y ahi hay variables, esas
	// variables van a persistir en cada llamado
	fmt.Println("Closures")
	a := incremento()
	b := incremento()
	c := incremento()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(c())
}

func incremento() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
