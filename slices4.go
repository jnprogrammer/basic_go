package main

import "fmt"

func main() {
	y := []int{2, 4, 8, 16, 32, 64, 128}
	fmt.Println(y)
	y = append(y, 256, 512, 1024, 2048)
	fmt.Println(y)

	x := []int{1, 3, 6, 12, 24, 48, 96}
	// ... means its a variadic parameter, means it takes an unlimited amount of arguments
	x = append(x, y...)
	fmt.Println(x)
}
