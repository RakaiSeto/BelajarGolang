package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"Name"`
	Age  int `json:"Age"`
}

func main() {
	s := `[{"Name":"Rakai Seto","Age":16},{"Name":"Joko Jiwo","Age":51}]`

	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var people []Person
	
	err := json.Unmarshal(bs, &people) 
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(people)
	
}