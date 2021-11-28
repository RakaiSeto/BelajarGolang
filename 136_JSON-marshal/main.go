package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := Person{
		Name: "Rakai Seto",
		Age:  16,
	}
	p2 := Person{
		Name: "Joko Jiwo",
		Age:  51,
	}

	people := []Person{
		p1, p2,
	}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}