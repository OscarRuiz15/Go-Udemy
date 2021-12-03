package main

import "fmt"

func main() {
	foo()
	bar("James")
	s := woo("Moneypenny")
	fmt.Println(s)
	x, y := saludar("Oscar", "Ruiz")
	fmt.Println(x, y)

	fmt.Println()
	sumi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Print(sumi)
	suma := sum("Oscar", sumi...)
	fmt.Println("El valor de la suma es ", suma)

	defer deferfoo()
	defer deferbar()
	nodefer()
}

//Receptor -> Valor que recibe una funcion, un valor que tiene adjuntada una funcion llamada metodo
//Cuando hay un receptor, es porque es un metodo
//func (r receptor) identificador(parametros) retorno(s) { codigo }
func foo() {
	fmt.Println("Hola desde foo.")
}

func bar(s string) {
	fmt.Println("Hola", s)
}

func woo(s string) string {
	return fmt.Sprint("Hola desde woo, ", s)
}

func saludar(n string, a string) (string, bool) {
	p := fmt.Sprint(n, " ", a, ` dice "hola"`)
	q := true

	return p, q
}

func sum(y string, x ...int) int {
	fmt.Println("\n", y, x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	suma := 0
	for _, v := range x {
		suma += v
	}
	fmt.Println(suma)
	return suma
}

func deferfoo() {
	fmt.Println("foo")
}

func deferbar() {
	fmt.Println("bar")
}

func nodefer() {
	fmt.Println("Jaja")
}