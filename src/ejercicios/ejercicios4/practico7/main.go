package main

import "fmt"

func main() {
	s1 := []string{"Batman", "Jefe", "Vestido negro"}
	s2 := []string{"Robin", "Ayudame", "Vestido de colores"}
	s := [][]string{s1, s2}
	fmt.Println(s)

	for key, value := range s {
		fmt.Println(key, value)
	}
	fmt.Println()
	for key, value := range s {
		fmt.Println(key, value)
		for key2, value2 := range value {
			fmt.Println("\t", key2, value2)
		}
		fmt.Println()
	}
}
