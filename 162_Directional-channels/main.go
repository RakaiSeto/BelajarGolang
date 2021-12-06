package main

import "fmt"

func main() {
	c := make(chan int, 1)
	cr := make(chan<- int)
	cs := make(<-chan int)

	c <- 42

	fmt.Println(c)

	fmt.Println("-----")
	fmt.Printf("%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cr\t%T\n", cs)

	// GENERAL (c) CAN CONVERT TO SPECIFIC (cs, cr), BUT NOT OTHERWISE
}