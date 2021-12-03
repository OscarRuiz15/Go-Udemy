package main

import "fmt"

func main() {
	opu := (16 == 15)
	opd := (17 <= 18)
	opt := (21 >= 5)
	opc := (4 != 9)
	opci := (4 < 6)
	ops := (14 > 66)
	fmt.Println(opu, opd, opt, opc, opci, ops)
}
