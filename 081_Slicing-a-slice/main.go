package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 37}
	fmt.Println(x[1])
	fmt.Println(x[1:])
	fmt.Println(x[1:4])

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}