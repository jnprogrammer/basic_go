//Using Iota create 4 constants for the last 4 years, print the contstant values
package main

import "fmt"

const (
	y1 = 2017 + iota
	y2
	y3
	y4
)

func main() {
	fmt.Println(y1)
	fmt.Println(y2)
	fmt.Println(y3)
	fmt.Println(y4)
}
