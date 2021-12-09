package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	salir := make(chan int)

	//enviar
	go enviar(par, impar, salir)

	//recibir
	recibir(par, impar, salir)

	fmt.Println("Finalizando")

}

func enviar(par, impar chan<- int, salir chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}

	//close(par)
	//close(impar)
	//salir <- 0
	close(salir)
}

func recibir(par, impar <-chan int, salir <-chan int) {
	for {
		select {
		case v := <-par:
			fmt.Println("Desde el canal par:", v, salir)
		case v := <-impar:
			fmt.Println("Desde el canal impar:", v, salir)
		case v, ok := <-salir:
			fmt.Println("Desde el canal salir:", v, ok)
			return
		}
	}
}
