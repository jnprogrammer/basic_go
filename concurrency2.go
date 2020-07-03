package main

import "fmt"

func roo(x int) int {
	return x * 420
}
func main() {
	ch := make(chan int)

	go func() {
		ch <- roo(710)
	}()
	fmt.Println(<-ch)
}

// Review the golang documentation always !!
