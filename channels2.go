package main

import "fmt"

func main() {
	//c := make(chan <- int, 2) //directional channel send only
	//c := make(<- chan int, 2) //directional channel receive only
	c := make(chan int)
	creceive := make(<-chan int)
	csend := make(chan<- int)
	go func() {
		c <- 423
	}()

	fmt.Println(<-c)
	fmt.Printf("%T\n", c)
	fmt.Println("Receive Channel", creceive)
	fmt.Println("Send Channel", csend)

}
