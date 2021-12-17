package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Rakai")
	if s != "Welcome,Rakai" {
		t.Error("got", s, "expected", "Welcome,Rakai")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Rakai"))
	// Output:
	// Welcome,Rakai
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Rakai")
	}
}