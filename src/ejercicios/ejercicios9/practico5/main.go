package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var incremento int64
	const gs = 50
	wg := sync.WaitGroup{}
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incremento, 1)
			fmt.Println(atomic.LoadInt64(&incremento))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Cuenta:", incremento)
}
