package main

import "fmt"

func main() {
	c := make(chan int)

	// Send
	go foo(c)
	// Receive
	bar(c)

	fmt.Println("About to exit")
}

// Send
func foo(c chan<- int) {
	c <- 42
}

// Receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}