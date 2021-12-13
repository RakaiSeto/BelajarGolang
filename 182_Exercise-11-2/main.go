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

	bs, err := toJSON(p1)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(bs))
}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("it's an error bro, %v", err)
	}
	return bs, nil
}