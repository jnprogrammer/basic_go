package main

import "fmt"

var x int

func main() {
	fmt.Println(x)
	x++
	fmt.Println(x)
	{
		x := 534
		fmt.Println("This x is in a different scope:", x)
	}
	foo()
	fmt.Println(x)

	a := incrementor()
	b := incrementor()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func foo() {
	fmt.Println("Scopre of x is the entire package")
	x++
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
