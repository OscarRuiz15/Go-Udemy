package main

import "fmt"

//Errores con panic
func main() {
	var resp1, resp2, resp3 string

	fmt.Println("Nombre")
	_, err := fmt.Scan(&resp1)
	if err != nil {
		panic(err)
	}

	fmt.Println("Comida favorita")
	_, err = fmt.Scan(&resp2)
	if err != nil {
		panic(err)
	}

	fmt.Println("Deporte favorito")
	_, err = fmt.Scan(&resp3)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp1, resp2, resp3)
}
