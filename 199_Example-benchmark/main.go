package main

import (
	"fmt"
	"strings"

	mystr "github.com/RakaiSeto/belajarGolang/199_Example-benchmark/myStr"
)

const s = "Exercitation enim qui adipisicing fugiat do. Consectetur ad do ea nostrud mollit irure ex id aliquip irure in cupidatat ut. Lorem consectetur do commodo nisi voluptate reprehenderit proident reprehenderit ad. Sit labore deserunt ad consectetur laborum minim in laboris."

func main() {
	xs := strings.Split(s, " ")

	for _, v := range xs {
		fmt.Println(v)
	}

	fmt.Printf("\n%s\n", mystr.Cat(xs))
	fmt.Printf("\n%s\n\n", mystr.Join(xs))
}