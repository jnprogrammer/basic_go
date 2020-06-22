package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Println(x)

	// This is how to slice a slice

	//Starts at index 1 not including it, prints out the rest of the slice
	fmt.Println(x[1:])
	//Starts at index 1 not including it, prints up to index 4 but doesn't include index 4 in what it prints
	fmt.Println(x[0:4])
}
