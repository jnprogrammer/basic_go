package main

import (
	"testing"
)

func TestMySum(t *testing.T) {
	if mySum(2, 3) != 5 {
		t.Errorf("Expected", 5, "Got", x)
	}
}
