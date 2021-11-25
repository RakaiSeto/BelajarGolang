package main

import "fmt"

func main() {
	numSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	s := sum(numSlice...)
	fmt.Println(s)
	
	s2 := even(sum, numSlice...)
	fmt.Println(s2)
	
	s3 := odd(sum, numSlice...)
	fmt.Println(s3)
	
}

func sum(toSum ...int) int {
	total := 0
	for _, v := range toSum {
		total += v
	}
	return total
}

func even (f func(x ...int) int, toBeEven ...int) int {
	var evNum []int
	for _, v := range toBeEven {
		if v % 2 == 0 {
			evNum = append(evNum, v)
		}
	}
	return f(evNum...)
}

func odd (f func(x ...int) int, toBeOdd ...int) int {
	var oddNum []int
	for _, v := range toBeOdd{
		if v % 2 != 0{
			oddNum = append(oddNum, v)
		}
	}
	return f(oddNum...)
}