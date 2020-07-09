package main

import (
	"fmt"
	"sync"
)

func main() {
	evench := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	//send
	go send3(evench, odd)

	//receive
	go receive3(evench, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

//				 receive channel	 send channel
func receive3(e, o <-chan int, fanin chan<- int) {
	var waitersgroup sync.WaitGroup
	waitersgroup.Add(2)

	go func() {
		for v := range e {
			fanin <- v
		}
		waitersgroup.Done()
	}()

	go func() {
		for v := range o {
			fanin <- v
		}
		waitersgroup.Done()
	}()

	waitersgroup.Wait()
	close(fanin)
}

func send3(e, o chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

//Remember, Channels Block
//comma okay idiom can be used to check if a channel is closed
