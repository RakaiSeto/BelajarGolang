package main

import "fmt"

func main() {
	m := map[string]int{
		"Rakai": 16,
		"Joko":  51,
		"Sese":  49,
	}
	fmt.Println(m)
	fmt.Println(m["Rakai"])
	fmt.Println(m["Joko"])
	fmt.Println(m["Sese"])

	fmt.Println(m["tes"]) //gada orang ini, tapi tetep keluar

	//cara ngecek
	v, ok := m["tes"]
	fmt.Println(v)
	fmt.Println(ok)
	if v, ok := m["tes"]; ok {
		fmt.Println("ini tes doang", v)
	}
	if v, ok := m["Rakai"]; ok {
		fmt.Println("ini tes doang", v)
	}
}