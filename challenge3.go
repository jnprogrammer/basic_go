package main

import (
	"fmt"
)

//  := short declaration operator creates a new variable
type grinder int //making own type with a underlying type of int
type bong int

var x grinder
var y bong
var z int

func main() {
	fmt.Printf("%T\n", x)
	fmt.Println(x)
	x = 710
	fmt.Println(x)

	z = int(x)
	fmt.Printf("%T", y)
	fmt.Println(y)

}

//Made my own type with an underlying type of int, then printed the type of the var with the
// identifier x then assigned an into x and printed x.
