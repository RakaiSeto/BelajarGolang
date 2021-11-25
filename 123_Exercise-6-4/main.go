package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p := person{
		first: "Rakai",
		last:  "Seto",
		age:   16,
	}
	p.speak()
}

func (p person) speak() {
	fmt.Println("my name is", p.first, p.last, "and my age is", p.age)
}