package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func main() {
	c := circle{
		radius: 7,
	}
	info(&c)
}

// Area Func
func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface{
	area() float64
}

func info (s shape) {
	fmt.Println("The area is", s.area())
}