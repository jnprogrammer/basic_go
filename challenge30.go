package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100}
	f1 := foo(x...)
	fmt.Println(f1)
	fmt.Printf("%v", f1)

	f2 := boo(x)
	fmt.Println("this is f2 ", f2)
}

func foo(i ...int) int {
	num := 0
	for _, v := range i {
		num += v
	}
	return num
}

func boo(i []int) int {
	num := 0
	for _, v := range i {
		num += v
	}
	return num
}
