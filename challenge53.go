//Fix the dead lock

//c := make(chan int)
//
//c <-42
//
//fmt.println(<-c)

package main

import "fmt"

func main() {

	c := make(chan int, 32)
	//go func() {
	//	c <-42
	//}()
	c <- 42
	fmt.Println(<-c)
}
