package main

import (
	"fmt"
)

//const(
//	a = iota + 3124
//	b = iota
//	c = iota
//)

// TODO: iotoa will increment all var in a const and reset when the next const is declared.

//const(
//	a = iota
//	b
//	c
//)
//
//
//const(
//	e = iota
//	f
//	g
//)

const (
	a = iota
	b
	c
	e
	f
	g
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

}
