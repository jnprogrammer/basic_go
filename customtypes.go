package main

import (
	"fmt"
)

var a int

type hotwings int
type salsa string

var b hotwings
var c salsa

func main() {
	a = 420
	b = 710
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	c = "Hot wings"

	fmt.Println(b)
	fmt.Printf("%T\n", b)

	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = hotwings(a)
	fmt.Println(c)
	fmt.Printf("%T\n", b)

}
