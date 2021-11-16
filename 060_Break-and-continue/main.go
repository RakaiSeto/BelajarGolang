package main

import "fmt"

func main() {
	x := 41 / 40
	fmt.Println(x)

	even()
}

func even() {
	x := 0
	for {
		x++
		if x > 100 {
			break		
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
}