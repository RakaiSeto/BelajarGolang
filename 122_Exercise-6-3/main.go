package main

import "fmt"

func main() {
	defer deffered()
	foo()
	bar()
}

func deffered() {
	fmt.Println("1")
}
func foo() {
	fmt.Println("2")
}
func bar() {
	fmt.Println("3")
}