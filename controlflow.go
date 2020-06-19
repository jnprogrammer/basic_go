package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	//   init ; condition; execute after 1 trip of the for loop
	for x := 0; x <= 420; x++ {
		fmt.Println(x)
	}
}
