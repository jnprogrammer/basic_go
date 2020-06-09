package main

import (
	"fmt"
)

var a int

type hotwings int

var b hotwings

func main() {
	a = 420
	b = 710
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(b)
	fmt.Printf("%T\n", b)

	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = hotwings(a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

}
