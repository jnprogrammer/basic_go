package main

import "fmt"

func main() {

	c := make(chan int)

	//sends
	go func() {
		for i := 0; i < 100; i++ {
			c <- 42 + i
		}
		close(c)
	}()
	//receive
	for v := range c { //range will keep looping over a channel till its closed.
		fmt.Println(v)
	}
}

//func foo(c chan<- int) { //you can go from general to specific.
//	for i := 0; i < 100; i++{
//		c <- 42 + i
//	}
//	close(c) //closes the channel
//}
