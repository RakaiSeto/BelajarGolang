package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("GO ROUTINES", runtime.NumGoroutine())
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
	}
	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	fmt.Println("Count:", counter)
}