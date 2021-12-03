package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("Número de CPU:", runtime.NumCPU())
	fmt.Println("Número de GoRutines:", runtime.NumGoroutine())

	var contador int64
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	//var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			//mu.Lock()
			//v := contador
			//runtime.Gosched()
			//v++
			//contador = v
			//mu.Unlock()
			fmt.Println("Contador:", atomic.LoadInt64(&contador))
			wg.Done()
		}()
		fmt.Println("Número de GoRutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Cuenta:", contador)
}
