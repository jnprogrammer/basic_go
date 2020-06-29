package main

import (
	"fmt"
)

func main() {
	fmt.Println(4 * 3 * 2 * 1)
	n := fact2(56)
	fmt.Println(n)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func fact2(number int) int {
	total := 1
	for ; number > 0; number-- {
		total *= number
	}
	return total
}
