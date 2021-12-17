package main

import (
	"fmt"
	"Github.com/RakaiSeto/BelajarGolang/192_Exercise-12-1/dog"
)

type anjing struct {
	name string
	age int
}

func main() {
	monmon := anjing{"monmon", dog.DogYear(6)}
	fmt.Printf("Nama anjing: %v \nUmur anjing: %v", monmon.name, monmon.age)
}

