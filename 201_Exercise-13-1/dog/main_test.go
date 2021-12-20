package dog

import (
	"testing"
	"fmt"
)

func TestDogYear(t *testing.T) {
	x := DogYear(2)
	if x != 14 {
		t.Error("Expected 14, got", x)
	}
}

func ExampleDogYear() {
	fmt.Println(DogYear(2))
	// Output:
	// 14
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DogYear(2)
	}
}
