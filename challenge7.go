//Assign an int to a variable
//Print that int in decimal, binary, and hex
//shift the bits of the int over 1 and assign it to the variable
//print that variable again in decimal, binary and hex

package main

import (
	"fmt"
)

var number int

func main() {
	number = 710

	fmt.Printf("Hex %X \nBinary %b \nDecimal %v\n", number, number, number)
	number = number << 1
	fmt.Print("Number was bitshifted << 1\n")
	fmt.Printf("Hex %X \nBinary %b \nDecimal %v\n", number, number, number)
	number = number >> 2
	fmt.Print("Number was bitshifted >> 2\n")
	fmt.Printf("Hex %X \nBinary %b \nDecimal %v\n", number, number, number)
	number2 := number * 2
	fmt.Printf("Hex %X \nBinary %b \nDecimal %v\n", number2, number2, number2)

}
