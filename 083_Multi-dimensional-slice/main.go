package main

import "fmt"

func main() {
	rakai := []string{"Rakai", "Seto", "Coklat", "Moka"}
	joko := []string{"Joko", "Jiwo", "Duren", "Kopi"}
	fmt.Println(rakai)
	fmt.Println(joko)
	
	xp := [][]string{rakai, joko}
	fmt.Println(xp)
}