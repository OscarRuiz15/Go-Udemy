package main

import "fmt"

func main() {
	//LITERAL COMPUESTO (COMPOSITE LITERAL)
	x := []int{2, 3, 4, 42, 8, 72}
	fmt.Println(x, len(x))
	fmt.Printf("%T\n", x)
	fmt.Println()
	for key, value := range x {
		fmt.Println(key, value)
	}
	fmt.Println()
	for _, value := range x {
		fmt.Println(value)
	}

	//SLICING
	fmt.Println()
	fmt.Println(x[1:5]) //Desde donde, hasta donde -1
	fmt.Println()
	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}
	//Agregar elementos a un slice
	fmt.Println()
	x = append(x, 44, 55, 66)
	fmt.Println(x, len(x))
	y := []int{99, 98, 97}
	x = append(x, y...)
	fmt.Println(x, len(x))

	//Remover elementos a un slice
	fmt.Println()
	x = append(x[:4], x[7:]...)
	fmt.Println(x, len(x))

	//Funcion make
	fmt.Println("\nMake")
	x = make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 1
	x[5] = 100
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 200)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 300)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println("Se llena")
	x = append(x, 400)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	//Slices multidimensionales
	fmt.Println()
	et := []string{"Oscar", "Ruiz", "Crossfit", "Baseball", "Montañismo"}
	fmt.Println(et)
	j1 := []string{"Jacinto", "Lopez", "Correr", "Nadar", "Bailar"}
	fmt.Println(j1)

	vp := [][]string{et, j1}
	fmt.Println(vp)
	fmt.Println(vp[0][3])
	vp[1][2] = "Modifica"
	fmt.Println(vp)

	//Slice - Arreglo subyacente
	fmt.Println()
	x = []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	y = append(x[:2], x[3:]...) //Se utiliza el mismo arreglo subyacente
	fmt.Println(y)
	fmt.Println(x)
	fmt.Println()

	x = []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	y = append(x, 6, 7, 8) //Un nuevo arreglo subyacente es asignado
	fmt.Println(y)
	fmt.Println(x)
}
