package main

import "fmt"

func main() {
	foo()
	bar("Tera)")
	wrds := retro("420")
	fmt.Println(wrds)

	x, y := dude("A ", "Dude")

	fmt.Println(x)
	fmt.Println(y)
}

// func (r receiver) identifier(paramerters) (returns(s)) {  ...  }

func foo() {
	fmt.Println("Hey Go functions")
}

func bar(s string) {
	fmt.Println("Hey ", s)
}

func retro(s string) string {
	return fmt.Sprint("Hello hey im printing to a string with sPrint or string print ", s)
}

// Everything in Go is PAss by Value

func dude(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, `,says hey`)

	b := false
	return a, b
}
