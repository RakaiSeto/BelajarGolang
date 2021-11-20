package main

import "fmt"

func main() {
	m := map[string]int{
		"Rakai": 16,
		"Joko":  51,
		"Sese":  49,
	}
	fmt.Println(m)
	
	delete(m, "Rakai")
	fmt.Println(m)
}