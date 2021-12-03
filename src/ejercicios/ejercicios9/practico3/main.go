package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var incremento int64
	const gs = 50
	wg := sync.WaitGroup{}
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			val := incremento
			runtime.Gosched()
			val++
			incremento = val
			fmt.Println(incremento)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Cuenta:", incremento)
}
