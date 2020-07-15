package main

import (
	"testing"
)

func TestMySum(t *testing.T) {
	x := mySum(2, 3)
	if x != 5 {
		t.Errorf("Expected", 5, "Got", x)
	}
}
