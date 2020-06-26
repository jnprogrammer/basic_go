package main

import "fmt"

func main() {
	//sli := []int{
	//	2,4,6,8,10,12,14,
	//}
	// I unfurled the slice in the
	sum("tst", 532, 45, 6)

}

//functions can take zero or more arguments
func sum(s string, x ...int) int {
	var number int
	for i, v := range x {
		number += v
		fmt.Println(i, number)
	}
	return number
}
