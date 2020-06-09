package main

import (
	"fmt"
)

var y string //when a var is declared without a initial value it is set to zero

func main() {
	fmt.Println("Value of y is: ", y)
	fmt.Printf("%T\n", y)

	y = "Bond, James Bond"

}
