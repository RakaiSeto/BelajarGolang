package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 3, 32, 42, 99, 18, 16, 56, 12}
	xs := []string{"Rakai", "Joko", "Ariq", "Rili", "Nala", "Sya"}

	fmt.Println(xi)
	fmt.Println(xs)

	sort.Ints(xi)
	sort.Strings(xs)

	fmt.Println("------")
	fmt.Println(xi)
	fmt.Println(xs)
}
