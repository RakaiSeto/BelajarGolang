package main

import "fmt"

func main() {
	m := map[string][]string{
		"rakai": {"tes", "tes2", "tes3"},
		"joko":  {"pak'e", "tes4", "tes5"},
	}

	fmt.Println(m)

	m["sese"] = []string{"roti", "kue", "bolu"}

	for k, v := range m{
		fmt.Println("This is the record", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}