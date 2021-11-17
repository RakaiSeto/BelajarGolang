package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this aint gonna print")
	case (2 == 4):
		fmt.Println("this aint gonna print either")
	case (4 != 3):
		fmt.Println("this gon print")
	case (3 != 4):
		fmt.Println("although true, aint gon print")
	}

	switch {
	case false:
		fmt.Println("this aint gonna print")
	case (2 == 4):
		fmt.Println("this aint gonna print either")
	case (6 != 5):
		fmt.Println("this gon print")
		fallthrough
	case (5 != 6):
		fmt.Println("with fallthrough, its gon print")
	}

	switch {
	case false:
		fmt.Println("this aint gonna print")
	case (2 == 4):
		fmt.Println("this aint gonna print either")
	case (6 == 5):
		fmt.Println("this gon print")
	case (5 == 6):
		fmt.Println("with fallthrough, its gon print")
	default:
		fmt.Println("when nothing clicks, its gon print 'default'")
	}
}