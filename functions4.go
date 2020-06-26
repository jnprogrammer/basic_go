package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("fo")
}

func bar() {
	fmt.Println("bar")

}
