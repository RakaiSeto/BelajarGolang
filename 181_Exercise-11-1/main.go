package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		"James",
		"Bond",
		[]string{"Bla bla bla", "punten kakanya", "monggo adenya"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
}