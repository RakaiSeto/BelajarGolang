package main

import "fmt"

type person struct{
	name string
	age int
}

func main() {
	x := person{
		name: "Rakai Seto",
		age: 16,
	}
	fmt.Println(x)
	changeMe(&x)
	fmt.Println(x)
}

func changeMe (p *person) {
	p.name = "bruh"
}