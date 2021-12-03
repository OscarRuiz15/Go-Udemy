package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `clave hackeable`
	fmt.Println([]byte(s))
	bs, err := bcrypt.GenerateFromPassword([]byte(s), 4)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(s)
	fmt.Println(bs)
	fmt.Println(string(bs))

	claveLogin := `clave hackeable`
	fmt.Println([]byte(claveLogin))
	err = bcrypt.CompareHashAndPassword(bs, []byte(claveLogin))
	if err != nil {
		fmt.Println("No logea", err)
		return
	}
	fmt.Println("Logea")
}
