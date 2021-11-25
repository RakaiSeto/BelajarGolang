package main

import "fmt"

func main() {
	fmt.Println(foo(16))
	fmt.Println(bar(16))
}

func foo(x int) int {
	x++
	return x
}

func bar(x int) (string, int) {
	x--
	return "x decrement result is", x
}