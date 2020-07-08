package main

import "fmt"

func main() {

	c := make(chan int)

	//sends
	go foo(c)

	//receive
	bar(c) //if I add go here the program will exit to fast for the channel to receive anything
	fmt.Println("About to exit")
}

// func send
func foo(c chan<- int) { //you can go from general to specific.
	c <- 438
}

// func receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
