package main

import "fmt"

func main() {
	// func (r receiver) identifier(parameters) (return(s)) {code}

	foo()
	bar("Rakai")
	s1 := woo("Rakai")
	fmt.Println(s1)
	x, y := mouse("Rakai", "Seto")
	fmt.Println(x)
	fmt.Println(y)
}

func foo() {
	fmt.Println("Hello from foo")
}

func bar(s string)  {
	fmt.Println("Hello,", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from Woo, ", s)
}

func mouse(f, l string) (string, bool) {
	s := fmt.Sprint(f, " ", l, `, says "Hello"`)
	a := false
	return s, a
}