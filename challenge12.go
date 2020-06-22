// print a for loop using the {} style
package main

import "fmt"

func main() {
	year := 1776
	count := 0

	for year <= 2020 {
		count++
		fmt.Println(year)
		year++
	}
	fmt.Println(count)
}
