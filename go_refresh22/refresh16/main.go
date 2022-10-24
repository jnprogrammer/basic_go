package main

import (
	"fmt"
)

func main() {
	//slices don't store data, they point to an array
	course := make([]string, 5, 10)
	course[0] = "Reading books"
	course[1] = "Baking books"
	course[2] = "Eating books"

	fmt.Printf("Length of slice is %d and capacity is %d\n",
		len(course), cap(course))

	fmt.Println(course)
}
