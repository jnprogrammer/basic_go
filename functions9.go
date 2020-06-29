package main

import "fmt"

func main() {
	f1 := foo()
	fmt.Println(f1)

	//b1 := func() int {
	//	return 420
	//}

	//b2 := bars()

	//s := b2()
	fmt.Println(bars()()) //s // each () is a level of a function

	//fmt.Printf("%T\n" ,b1)
	fmt.Println("The number returned from the function that was returned ")
}

func foo() string {

	return "Hey all, I'm a function that is returning a this string"
}

func bars() func() int {
	return func() int {
		return 420
	}
}
