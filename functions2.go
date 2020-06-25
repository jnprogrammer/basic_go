package main

import "fmt"

func main() {
	foo(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 820)
}

func foo(x ...int) {
	fmt.Println("Stuff in text")
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	for range x {
		fmt.Println(x)
	}
}

// func (r receiver) identifier(parameters(s)) (return(s)) { code }
