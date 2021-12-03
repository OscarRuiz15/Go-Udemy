package main

import (
	"fmt"
	"sync"
)

func main() {

	var incremento int64
	const gs = 50
	wg := sync.WaitGroup{}
	wg.Add(gs)
	sm := sync.Mutex{}

	for i := 0; i < gs; i++ {
		go func() {
			sm.Lock()
			defer sm.Unlock()
			val := incremento
			val++
			incremento = val
			fmt.Println(incremento)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Cuenta:", incremento)
}
