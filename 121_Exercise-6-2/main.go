package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a := foo(s...)
	b := bar(s)
	fmt.Println(a)
	fmt.Println(b)
}

func foo(x ...int) int {
	var total int
	for _, v := range x {
		total += v
	}
	return total
}

func bar(x []int) int {
	var total int
	for _, v := range x {
		total += v
	}
	return total
}