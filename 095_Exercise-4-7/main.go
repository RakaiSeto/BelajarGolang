package main

import "fmt"

func main() {
	rakai := []string{"rakai", "seto", "punten"}
	joko := []string{"joko", "jiwo", "halo"}
	xy := [][]string{rakai, joko}
	for i, xxy := range xy{
		fmt.Println("record:", i)
		for j, v := range xxy {
			fmt.Printf("\t%v\t%v\n", j, v )
		}
	}
}