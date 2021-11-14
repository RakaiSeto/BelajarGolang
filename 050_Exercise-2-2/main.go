package main

import "fmt"

func main() {
	g := 37 == 42
	h := 37 <= 42
	i := 37 >= 42
	j := 37 != 42
	k := 37 < 42
	l := 37 > 42

	fmt.Println(g, h, i, j, k, l)
}