package main

import "fmt"

func main() {
	n, err := fmt.Println("Hey you")

	if err == nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
