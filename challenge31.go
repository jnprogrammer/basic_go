package main

import "fmt"

func main() {

	foo()
	defer boo()
	roo()
}

func foo() {
	fmt.Println("This is foo, it comes now !!")
}

func boo() {
	fmt.Println("This is boo, it's coming after derfer allows it, which is after the block ends")
}

func roo() {
	defer func() {
		fmt.Println("roo DEFER ran")
	}()
	fmt.Println("roo ran")
}
