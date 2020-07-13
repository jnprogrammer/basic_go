package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
				fmt.Println("WHOA") // I wonder why this happens when it runs?
			}
		}()
	}

	for v := 0; v < 100; v++ {
		fmt.Println(v, <-c)
	}

}
