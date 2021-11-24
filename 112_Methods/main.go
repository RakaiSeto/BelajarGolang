package main

import "fmt"

type person struct {
	first string
	last  string
}

type pelajar struct {
	person
	sma bool
}

func (p pelajar) speak() {
	fmt.Println("I am", p.first, p.last)
}

func main() {
	p1 := pelajar{
		person: person{
			"Rakai",
			"Seto",
		},
		sma: true,
	}
	fmt.Println(p1)
	p1.speak()
}