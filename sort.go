package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []int{4, 5, 67, 2, 2, 5, 1, 3, 56, 64, 7, 234, 54, 6, 532, 6, 352, 5634, 6, 345, 345, 43, 566, 2345, 1234, 5432, 543, 253, 4543, 5, 234, 5, 6763}
	y := []string{"Tommny", "Fist", "s", "grub", "S"}
	fmt.Println(x)
	sort.Ints(x)
	fmt.Println(x)
	fmt.Println(y)
	sort.Strings(y)
	fmt.Println(y)
}
