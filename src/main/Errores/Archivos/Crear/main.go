package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

//Crear y escribir en archivo
func main() {
	f, err := os.Create("nombres.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("Oscar Ruiz")
	io.Copy(f, r)
}
