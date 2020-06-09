package main

import "fmt"

func main() {

	n, err := fmt.Println("Fires of Liberation")
	fmt.Println(n)
	fmt.Println(err)
	foo()
	bar()
}

func foo() {
	fmt.Println("I'm not foo")
}

func bar() {
	fmt.Println("THis is the end of the program")
}
