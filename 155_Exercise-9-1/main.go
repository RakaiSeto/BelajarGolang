package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	fmt.Println("sumn")
	go foo(4)
	go bar(5)
	wg.Wait()
	fmt.Println(runtime.NumGoroutine())
}

func foo(x int) {
	fmt.Println(x*x)
	wg.Done()
}

func bar(x int) {
	fmt.Println(x*3)
	wg.Done() 
}