package main

import (
	"testing"
)

func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}
	tests := []test{
		test{[]int{21, 41}, 42},
		test{[]int{1, 1}, 2},
		test{[]int{21, 10}, 31},
		test{[]int{400, 20}, 420},
		test{[]int{400, 20, -840}, 420},
	}
	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Errorf("Expected", v.answer, "Got", x)
		}
	}

}
