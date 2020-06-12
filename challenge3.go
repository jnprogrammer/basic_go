package main

import (
	"fmt"
)

type grinder int //making own type with a underlying type of int

var x grinder

func main() {
	fmt.Printf("%T\n", x)
	fmt.Println(x)
	x = 710
	fmt.Println(x)
}

//Made my own type with an underlying type of int, then printed the type of the var with the
// identifier x then assigned an into x and printed x.
