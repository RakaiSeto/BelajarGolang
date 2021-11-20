package main

import "fmt"

func main() {
	m := map[string]int{
		"Rakai": 16,
		"Joko":  51,
	}
	m["Sese"] = 49

	for k, v := range m{
		fmt.Println(k, v)
	}
}