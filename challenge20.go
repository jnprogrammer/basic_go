// make a slice, then append a value, then print out the new slice, then append 3 values in one statement to the slice
// then print out the slice again, then append that slice to a new slice.
package main

import "fmt"

func main() {
	x := []int{42, 53, 563, 53, 0, 54, 895}
	y := []int{8080, 2004, 403, 2025}
	fmt.Println("Original slice\n", x)
	fmt.Println("**********************\n")

	x = append(x, 420)
	fmt.Println(x)

	x = append(x, 710, 1130, 024)
	fmt.Println(x)

	y = append(y, x...)
	fmt.Println(y)
}
