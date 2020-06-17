package main

import (
	"fmt"
)

var x = 3

func main() {
	x := 4
	fmt.Printf("%d\t\t%b", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b", y, y)

	z := x << 45
	fmt.Printf("%d\t\t%b", z, z)

}
