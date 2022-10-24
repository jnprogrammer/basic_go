package main

import (
	"fmt"
)

func main() {
	//slices don't store data, they point to an array
	course := []string{"Reading books", "Baking books", "Eating books"}

	fmt.Printf("Length of slice is %d and capacity is %d\n",
		len(course), cap(course))

	fmt.Println(course)
}
