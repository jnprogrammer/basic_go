package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 7, 8, 9, 10, 11, 12}
	y := []int{5, 6, 7, 8, 12}

	fmt.Println(x, y)
	// Take a slice of the slice you want to keep and append the rest of the slice to it. don't for gett the dots ... to get infinite args
	x = append(x[:3], x[4:]...)
	// start at index 3 and stop at index 4 which removes 7 from the slice
	fmt.Println(x, y)
}
