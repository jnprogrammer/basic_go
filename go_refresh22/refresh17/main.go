package main

import (
	"fmt"
)

func main() {

	mySlice := []int{1, 2, 3, 909, 710, 6, 7, 8, 9, 10}

	fmt.Println(mySlice[3])

	mySlice[1] = 420
	fmt.Println(mySlice)

	//slice a slice is gotastic
	//
	sliceOfSlice := mySlice[3:]
	fmt.Println(sliceOfSlice)
}
