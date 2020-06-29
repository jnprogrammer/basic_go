package main

import "fmt"

func main() {
	x := 0
	fmt.Println("x before", &x)
	fmt.Println("x before", x)
	fooly(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)

}

func fooly(y *int) {
	fmt.Println(y)
	fmt.Println(*y)
	*y = 43
	fmt.Println(y)
	fmt.Println(*y)

}

// you want to use pointers is when you have a large piece of data
// and you don't want to move that data around, just point to its address location
