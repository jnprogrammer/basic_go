package main

import "fmt"

func main() {
	c := make(chan int, 100) // that 100 is the buffer length.

	//go func(){
	//	c <- 420
	//}()
	c <- 710
	c <- 7168
	//c <-7186
	//c <-7780
	//c <-760
	//c <-720
	//c <-730
	//c <-740

	fmt.Println(<-c, <-c)
}

// Go channels work by a receive/send lock by channels, without this you'll get a deadlock
