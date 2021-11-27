package main

import "fmt"

func main() {
	x := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	s := sum(x...)
	fmt.Println(s)

	s2 := even(sum, x...)
	fmt.Println(s2)
}

func sum(x ...int) int {
	total := 0
	for _, v := range x{
		total += v
	}
	return total
}

func even (f func(x ...int) int, toEven ...int) int {
	var even []int
	for _, v := range toEven {
		if v%2 == 0 {
			even = append(even, v)
		}
	} 
	return f(even...)
}