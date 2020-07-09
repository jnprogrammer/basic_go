package main

import "fmt"

func main() {
	evench := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	//send
	go send2(evench, odd, quit)

	//receive
	receive2(evench, odd, quit)

	fmt.Println("about to exit")
}

func receive2(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("From the EVEN !! channel: ", v)
		case v := <-o:
			fmt.Println("From the ODD??   channel: ", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("From comma ok ", i, ok)
				return
			} else {
				fmt.Println("from comma ok ", i, ok)
			}
		}
	}
}

func send2(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	//close(e)
	//close(o)
	close(q)
}

//Remember, Channels Block
//comma okay idiom can be used to check if a channel is closed
