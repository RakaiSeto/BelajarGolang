package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 37, 38, 72, 94}
	fmt.Println(x)
	
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}