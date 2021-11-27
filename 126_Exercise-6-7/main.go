package main

import "fmt"

func main() {
	x := func() {
		fmt.Println("this is a func")
	}
	fmt.Println(x)
}