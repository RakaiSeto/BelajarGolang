package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type pelajar struct {
	person
	sekolah bool
}

func main() {
	p1 := pelajar{
		person: person{
			first: "Rakai",
			last:  "Seto",
			age:   16},
		sekolah: true,
	}
	p2 := person{
		first: "Joko",
		last:  "Jiwo",
		age:   51,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.person)
	fmt.Println(p1.person.first)
}
