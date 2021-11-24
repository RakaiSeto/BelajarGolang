package main

import "fmt"

func main() {
	s := struct {
		first   string
		last    string
		gf      bool
		age     int
		friends map[string]int
	}{
		first: "Rakai",
		last:  "Seto",
		gf:    true,
		age:   16,
		friends: map[string]int{
			"ariq": 123,
			"rili": 456,
			"nala": 789,
		},
	}

	fmt.Println(s)
	for k, v := range s.friends {
		fmt.Println(k, v)
	}
}
