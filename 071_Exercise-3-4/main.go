package main

import "fmt"

func main() {
	years := 2005
	for {
		if years > 2021 {
			break
		}
		fmt.Println(years)
		years++
	}
}