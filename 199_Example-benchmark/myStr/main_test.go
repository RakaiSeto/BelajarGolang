package mystr

import (
	"testing"
	"strings"
)

const s = "Exercitation enim qui adipisicing fugiat do. Consectetur ad do ea nostrud mollit irure ex id aliquip irure in cupidatat ut. Lorem consectetur do commodo nisi voluptate reprehenderit proident reprehenderit ad. Sit labore deserunt ad consectetur laborum minim in laboris."

var xs = strings.Split(s, " ")

func BenchmarkCat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cat(xs)
	}

}
func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join(xs)
	}
}