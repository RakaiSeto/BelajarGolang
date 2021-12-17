package main

import "testing"

func TestMySum(t *testing.T) {
	type test struct {
		data []int
		answer int
	}

	tests := []test{
		test{[]int{21, 21}, 42},
		test{[]int{3, 4, 5}, 12},
		test{[]int{1, 1}, 2},
		test{[]int{72, 13}, 85},
		test{[]int{16, 16}, 32},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("expected", v.answer, "got", x)
		}
	}

}