package main

import "fmt"

const (
	proxUno        = 2021 + iota
	proxDos        = 2021 + iota
	proxTres       = 2021 + iota
	proxCuatro     = 2021 + iota
)

func main() {
	fmt.Println(proxUno, proxDos, proxTres, proxCuatro)
}
