package main

import (
	"fmt"
)

var number int

func main() {
	number = 420

	fmt.Printf("In Binary %b\n", number)
	fmt.Printf("In Decimal %v\n", number)
	fmt.Printf("In Hexadecimal %X", number)

}
