package main

import "fmt"

func main() {

	x := 1
	for x <= 10 {
		fmt.Print(x)
		if x == 8 {
			fmt.Print("\n8 Ball broke the for loop")
			break
		}
		x++
	}
	fmt.Println("\ndone")

}
