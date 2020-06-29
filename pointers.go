package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) //shows the address where the VALUE is stored.\

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	fmt.Println(*b) // * means your dereferences a pointer which will give you the value stored at that address.
	fmt.Println(*&a)

	*b = 432 // I changed the VALUE of the VARIABLE of a

	fmt.Println(a)
}
