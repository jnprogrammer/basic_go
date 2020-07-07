package main

import "fmt"

func main() {

	c := make(chan int)
	creceive := make(<-chan int)
	csend := make(chan<- int)

	go func() { //without a anonymous function as a gorutine this code will deadlock
		csend <- 423
		creceive = c

	}()

	fmt.Println(<-creceive)

	fmt.Printf("%T\n", csend)

}

// this code is focusing on the differnt directions in channels and general to specific type converts. Specific to general type doesn't
