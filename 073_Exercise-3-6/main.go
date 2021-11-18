package main

import "fmt"

func main() {
	x := "Rakai"
	if x == "Rakai" {
		fmt.Printf("x is %v", x)
	} else if x == "Andi" {
		fmt.Println("x is Andi")
	} else {
		fmt.Println("x is not Rakai or Andi")
	}
}