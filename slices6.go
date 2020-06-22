package main

import "fmt"

// built in function make

func main() {
	//x := []int{4, 6, 7, 8, 45}
	// type, inital slice length,  max capacity of the slice (you can't append beyond this number)
	y := make([]int, 10, 101)
	y[0] = 32
	y[9] = 420

	// make prevents Go from having to deal with the issue of appending arrays by
	// making a new array with the old and new values. with make, you can bank a capacity of
	// how much you can append that way Go won't need to spend resources on building new arrays when append is used.
	// which is common in other languages when dealing with expanding arrays

	fmt.Println(y)
}
