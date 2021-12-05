package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) speak() {
	fmt.Println("bisa ges")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		"Rakai", 16,
	}
	fmt.Println("you CAN pass type *person to saySomething(&p)")
	fmt.Println("you CAN'T pass type person to saySomething(p)")
	saySomething(&p)
}
