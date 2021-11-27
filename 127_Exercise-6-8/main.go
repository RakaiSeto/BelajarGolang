package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x())
}

func foo() func() string {
	return func() string {
		return "returned func"
	}
}