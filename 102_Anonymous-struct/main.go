package main

import "fmt"

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Rakai",
		last:  "Seto",
		age:   16,
	}

	fmt.Println(p1)
}