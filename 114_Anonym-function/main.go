package main

import "fmt"

func main() {
	func(x int) {
		fmt.Println("Anonymous Func", x)
	}(37)
}