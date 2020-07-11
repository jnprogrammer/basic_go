//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	cs := make(chan<- int)
//
//	go func() {
//		cs <- 42
//	}()
//	fmt.Println(<-cs)
//
//	fmt.Printf("------\n")
//	fmt.Printf("cs\t%T\n", cs)
//}
//Fix code above

package main

import (
	"fmt"
)

func main() {

	cs := make(chan int)

	go func() {
		cs <- 42
	}()

	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}

//I figured to make it a bidirectional channel.
