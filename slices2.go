package main

import "fmt"

func main() {
	x := []int{42, 420, 71, 710, 1130}
	fmt.Println("The length of x:", len(x))
	fmt.Println("3rd element in the slice", x[2])

	for i, v := range x {
		fmt.Println(i, v)
	}

}
