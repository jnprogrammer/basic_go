package main

import (
	"fmt"
)

var x int
var y int
var z int
var b bool

func main() {
	x = 34
	y = 458
	z = -45
	b = false

	if x > y {
		fmt.Println("x is greater")
	} else {
		fmt.Println("y is greater")
	}
	if z < x {
		fmt.Println("z is greater")
	}

	if (z + y) <= (x * y) {
		print("what are you doing??\n")
	}

	if (-(z) + y*69) >= (x * y) {
		print("what are you doing AGAIN?\n")
	}

	if b != true {
		fmt.Println("Made b true")
		b = true
	}

	a := (43 == 5489)
	fmt.Println(a)
}
