package main

import "fmt"

// Anonymous func

func main() {
	foo()

	func() {
		fmt.Println("This is an Anonymous function")
	}()

	func(s string) {
		fmt.Println("This anonymous function ran with:", s)
	}("Anonymous call!")
}

func foo() {
	fmt.Println("Foo ran")
}
