package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 37}
	fmt.Println(x)
	x = append(x, 38, 72, 94)
	fmt.Println(x)
	y := []int{99, 104, 928}
	x = append(x, y...)
	fmt.Println(x)
}