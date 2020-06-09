package main

import (
	"fmt"
)

var outside = 210 //use var when declaring a variable outside the main function also known as package level scope

var z int

//Declare there isa variable with the identifier z and its type is int
func main() {
	x := 210
	fmt.Print("Variatic parm: ", x+outside)
}
