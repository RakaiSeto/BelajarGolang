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

	m := map[string]person{
		p1.lastName:	p1,
	}

	fmt.Println(p1.firstName, p1.lastName)

	for i, v := range p1.favIceCream{
		fmt.Println(i, v)
	}

	for k,  v := range m{
		fmt.Println(k)
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for i, val := range v.favIceCream{
			fmt.Println(i, val)
		}
	}
}