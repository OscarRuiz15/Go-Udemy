package main

import "fmt"

func main() {
	//Direccion a un slot de memoria donde hay guardado un valor de cualquier tipo
	a := 42
	fmt.Println(a)
	fmt.Println(&a) // Direccion en memoria de la variable a

	fmt.Printf("%T\n", a)
	//*int indica que es un puntero hacia un slot de memoria donde hay un valor de ese tipo
	fmt.Printf("%T\n", &a)

	var b *int = &a // Asigna la direeccion en memoria de la variable 'a' a la variable b
	fmt.Println(b)
	fmt.Println(*b)

	*b = 43
	fmt.Println(*b)

	//CUANDO USAR PUNTEROS
	//1. SI SE TIENE UNA VARIABLE O VALOR QUE INVOLUCRA UNA CANTIDAD GRANDE DE DATOS
	//2. CUANDO QUEREMOS CAMBIAR DIRECTAMENTE UN VALOR QUE ESTA ALMACENADO EN UNA DIRECCION DE MEMORIA
	fmt.Println()
	x := 0
	fmt.Println("X antes", x)
	fmt.Println("X antes", &x)
	fooPuntero(&x)
	fmt.Println("X despues", x)
	fmt.Println("X despues", &x)

	fmt.Println()

	//			Receptores       Valores
	//-----------------------------------------------
	//			  (t T)          T y *T
	//		      (t *T)          *T

}

func fooPuntero(y *int) {
	fmt.Println("y antes", y)
	fmt.Println("y antes", *y)
	*y = 42
	fmt.Println("y despues", y)
	fmt.Println("y despues", *y)
}
