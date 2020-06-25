package main

import "fmt"

func main() {
	foo(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 820)
	sum(4, 6, 2, 64, 5236, 34, 62345, 62345, 6, 23456, 1345, 63452, 56, 2345)
}

func foo(x ...int) {
	fmt.Println("Stuff in text")
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	for range x {
		fmt.Println(x)
	}
}

func sum(x ...int) {
	var number int
	args := 0
	for i, v := range x {
		number += v
		fmt.Println("This is i", i)
		args++
	}
	fmt.Println("The Sum of arguments", number, "There are", args, "Arguments")
}

// func (r receiver) identifier(parameters(s)) (return(s)) { code }
