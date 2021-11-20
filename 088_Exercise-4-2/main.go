package main

import "fmt"

func main() {
	x := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", x)
}