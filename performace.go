package main

import "fmt"

func main() {
	fmt.Println(mySum(200, 220))
	fmt.Println(mySum(50, 20))
	fmt.Println(mySum(02, 2220))
}

func mySum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
