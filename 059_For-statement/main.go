package main

import "fmt"

func main() {
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	while()
}

func while()  {
	y := 1
	for {
		if y > 9 {
			break
		}
		fmt.Println("its", y)
		y++
	}
	fmt.Println("done")
}