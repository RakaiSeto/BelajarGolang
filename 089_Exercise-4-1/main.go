package main

import "fmt"

func main() {
	x := [5]int{7, 5, 3, 19, 37}

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", x)
}