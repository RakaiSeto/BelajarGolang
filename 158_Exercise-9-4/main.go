package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	inc := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			x := inc
			runtime.Gosched()
			x++
			inc = x
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	fmt.Println("Count:", inc)
}