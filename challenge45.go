package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []int{4, 5, 67, 2, 2, 5, 763}
	y := []string{"test", "cat", "sssrs", "grub", "Sasdf"}
	fmt.Println(x)
	sort.Ints(x)
	fmt.Println(x)
	fmt.Println(y)
	sort.Strings(y)
	fmt.Println(y)
}
