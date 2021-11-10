package main

import "fmt"

var x = `kata gw, "test"`
var y = 42
var z = "test"

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}