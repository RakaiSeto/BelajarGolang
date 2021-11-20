package main

import "fmt"

func main() { 
	x := []int{4, 5, 7, 8, 37}
	fmt.Println(len(x))

	for i, v := range x {
		fmt.Println(i, v)
	}
}