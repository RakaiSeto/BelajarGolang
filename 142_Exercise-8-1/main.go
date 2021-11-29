package main

import (
	"fmt"
	"encoding/json"
)
type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{"Rakai", 32}
	u2 := user{"Joko", 51}
	u3 := user{"Sese", 49}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

}