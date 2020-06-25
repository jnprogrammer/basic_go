package main

import "fmt"

func main() {
	foo()
	bar("Tera)")
}

// func (r receiver) identifier(paramerters) (returns(s)) {  ...  }

func foo() {
	fmt.Println("Hey Go functions")
}

func bar(s string) {
	fmt.Println("Hey ", s)
}

// Everything in Go is PAss by Value
