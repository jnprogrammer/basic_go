// print a for loop using the for {} style
package main

import "fmt"

func main() {
	year := 1776
	count := 0

	for {
		count++
		fmt.Println(year)
		year++
		if year >= 2021 {
			break
		}
	}
	fmt.Println(count)
}
