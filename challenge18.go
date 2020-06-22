// Use a COMPOSITE LITERAL to make an Array that holds 5 VALUES of TYPE int
// assign Values to each index position
// range over array and print the values out.
// use format printing and print out the type of the array

package main

import "fmt"

func main() {
	comp := []int{45, 53, 2, 420, 4522}

	for i, v := range comp {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", comp)

}
