package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//Abrir archivo
func main() {
	f, err := os.Open("nombres.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bs))
}
