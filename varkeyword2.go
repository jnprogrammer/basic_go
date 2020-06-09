package main

import (
	"fmt"
)

var y = 22
var text string

// Go is a static programming language
// A var can only hold the TYPE it was made originally for

func main() {

	y = 422213
	fmt.Printf("%T: ", y)
	fmt.Println(y)

	fmt.Println(text)
	fmt.Printf("Type: %T:", text)
	fmt.Println("\n")

	text = "Now this var isn't empty"

	fmt.Printf("Type: %T:", text)
	fmt.Println(text)
}
