package main

import (
	"fmt"
)

var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)

	a := 5
	b := 45
	fmt.Println(a == b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a != b)
}
