package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	favIceCream []string
}

func main() {
	p1 := person{
		firstName:   "Rakai",
		lastName:    "Seto",
		favIceCream: []string{"Chocolate", "Durian",},
	}

	fmt.Println(p1.firstName, p1.lastName)

	for i, v := range p1.favIceCream{
		fmt.Println(i, v)
	}
}