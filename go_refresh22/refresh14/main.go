package main

import (
	"fmt"
)

func main() {

	listfood := []string{
		"cookies",
		"bread",
		"corn",
		"beets"}

	shoppinglist := []string{
		"cookies",
		"beets"}

	for _, i := range listfood {
		fmt.Println(i)
		for _, j := range shoppinglist {
			if i == j {
				fmt.Println("We found what's on the list")
			}
		}
	}
}
