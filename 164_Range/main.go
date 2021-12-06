package main

import "fmt"

func main() {
	c := make(chan int)

	// Send
	go func(){
		for i := 0; i < 100; i++ {
			c <- 42
		}
		close(c)
	}()
	// Receive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("About to exit")
}

// Send