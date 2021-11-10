package main

import "fmt"

var z int //this gonna be zeroed

func main() {
	x := 42
	fmt.Println(x)
	var y = 100
	fmt.Println(y)
	fmt.Println(z)
}