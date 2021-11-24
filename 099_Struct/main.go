package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Rakai",
		last:  "Seto",
		age:   16,
	}
	p2 := person{
		first: "Joko",
		last:  "Jiwo",
		age:   51,
	}

	fmt.Println(p1)
	fmt.Println(p1.first)
	fmt.Println(p2)
	fmt.Println(p2.first)
}