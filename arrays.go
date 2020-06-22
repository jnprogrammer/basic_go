package main

import "fmt"

func main() {
	// Arrays are zero based indexes
	// In Go, length is part of an arrays type
	// In Go don't use Arrays, uses Slices

	var x [5]int
	fmt.Println(x)
	x[3] = 4
	fmt.Println(x)
	fmt.Println(len(x))
}
