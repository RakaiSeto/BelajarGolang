package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(100)
	inc := 0
	for i := 0; i < 100; i++ {
		go func(){
			x := inc
			runtime.Gosched()
			x++
			inc = x
			fmt.Println(inc)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(inc)
}