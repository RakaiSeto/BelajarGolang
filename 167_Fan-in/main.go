package main

import (
	"sync"
	"fmt"
)

func main() {
	e := make(chan int)
	o := make(chan int)
	fanIn := make(chan int)

	go send(e, o)
	go receive(e, o, fanIn)

	for v := range fanIn{
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func send(e, o chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

func receive(e, o <-chan int, fanIn chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range e {
			fanIn <- v
		}
		wg.Done()
	}()
	
	go func() {	
		for v := range o {
			fanIn <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanIn)
}