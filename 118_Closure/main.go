package main

import "fmt"

func main() {
	var x int
	fmt.Println(x)
	x++

	// Untuk mengclosure (membatasi) scope sebuah variable
	// Kita membuat block didalam block
	{
		y := 16
		fmt.Println(y)
	}

	a := increment()
	b := increment()
	fmt.Println(a)
	fmt.Println(a)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(b)
	fmt.Println(b)

	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println("Hello")
}

func increment() func() int{
	var x int
	return func() int {
		x++
		return x
	}
}