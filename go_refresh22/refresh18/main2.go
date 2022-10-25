package main

import (
	"fmt"
)

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice)

	for _, i := range mySlice {
		fmt.Println("\n", i)
	}

	newSlice := []int{10, 20, 30}
	mySlice = append(mySlice, newSlice...)
	fmt.Printf("myslice has %d and new len is %d and new cap is %d\n", mySlice, len(mySlice), cap(mySlice))
}
