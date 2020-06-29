package main

import "fmt"

func main() {
	nums := []int{1, 2, 4, 5, 67, 7, 8, 9, 4, 5}
	s := sum(nums...)
	fmt.Print("\n", s)

	s2 := even(sum, nums...)
	fmt.Println("\n", s2)

	s3 := even(sum, nums...)
	fmt.Println("\n", s3)
}

func sum(x ...int) int {
	//fmt.Printf("%T\n", x)
	total := 0
	for _, val := range x {
		total += val
	}
	return total
}

func even(f func(x ...int) int, numbers ...int) int {
	var y []int
	for _, val := range numbers {
		if val%2 == 0 {
			y = append(y, val)
		}
	}
	return f(y...)
}

func odd(f func(x ...int) int, numbers ...int) int {
	var y []int
	for _, val := range numbers {
		if val%2 != 0 {
			y = append(y, val)
		}
	}
	return f(y...)
}
