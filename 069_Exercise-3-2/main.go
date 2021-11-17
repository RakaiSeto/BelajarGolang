package main

import "fmt"

func main() {
	fmt.Printf("Alphabet with ASCII code looped 3 times\n")
	for i := 65; i <= 90; i++ {
		fmt.Println("ASCII", i, "gonna print alphabet:")
		for j := 1; j < 4; j++ {
			fmt.Printf("\t\t%#U\n", i)
		}
	}
}