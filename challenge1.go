package main

import (
	"fmt"
)

var x int
var y bool

func main() {
	x = 42
	y = true
	z := "James Bond?"
	fmt.Println(x, y, z)
	fmt.Printf("%T\n", z)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
}
