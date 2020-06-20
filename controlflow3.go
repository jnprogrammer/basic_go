package main

import "fmt"

func main() {

	x := 1
	for x <= 10 {
		fmt.Print(x)
		if x == 8 {
			fmt.Print("\n8 Ball broke the for loop")
			break //stops the forloop its within and go to the next block
		}
		x++
	}
	fmt.Println("\ndone")

}
