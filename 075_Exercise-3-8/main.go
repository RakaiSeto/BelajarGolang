package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("shouldnt print")
	case true:
		fmt.Println("should print")
	}
}