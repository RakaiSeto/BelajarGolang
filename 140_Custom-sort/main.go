package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

	type ByAge []Person
	type ByName []Person

	func (a ByAge) Len() int { return len(a) }
	func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
	func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
	func (a ByName) Len() int { return len(a) }
	func (a ByName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
	func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func main() {
	p1 := Person{"Rakai", 16}
	p2 := Person{"Joko", 51}
	p3 := Person{"Sese", 49}

	people := []Person{p1, p2, p3}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)

}