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

type human interface{
	speak()
}

// Start
func (p pelajar) speak() {
	fmt.Println("I am pelajar", p.first, p.last)
}
func (p person) speak() {
	fmt.Println("I am person", p.first, p.last)
}

func bar(h human) {
	fmt.Println("I was passed into bar", h)
}

// Start Main
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

	bar(p1)
}