// create a program that shows the if statement in action

package main

import "fmt"

func main() {
	x := 5
	y := 10
	z := 20

	if x < y {
		fmt.Println("y is greater than x")
	} else if y > x {
		fmt.Println("x is less than y")
	}
	if z > x {
		fmt.Println("z is greater than x")
	}

}
