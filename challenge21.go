//delete 3 values from a slice

package main

import "fmt"

func main() {
	// deletes all values at index 1 up to 5
	x := []int{32, 543, 566, 72, 198, 420}
	fmt.Println(x)
	x = append(x[:1], x[5:]...)
	fmt.Println(x)
}
