package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally form f")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoverted in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0) //modify this argument to change the behavior and see the lessons point.
	fmt.Println("Return normally from G")
}

func g(i int) {
	if i > 3 {
		fmt.Println("PANICKING!~!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
