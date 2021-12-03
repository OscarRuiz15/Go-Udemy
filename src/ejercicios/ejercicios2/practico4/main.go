package main

import "fmt"

func main() {
	entero := 42
	fmt.Println("Hexadecimal\tBinario\tDecimal")
	fmt.Printf("%#x\t\t%b\t%d\n", entero, entero, entero)
	entero = entero << 1
	fmt.Printf("%#x\t\t%b\t%d", entero, entero, entero)
}
