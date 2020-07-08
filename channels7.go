package main

import "fmt"

func main() {
	evench := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send
	go send(evench, odd, quit)

	//receive
	receive(evench, odd, quit)

	fmt.Println("about to exit")
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("From the EVEN !! channel: ", v)
		case v := <-o:
			fmt.Println("From the ODD??   channel: ", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("From comma ok ", i)
				return
			} else {
				fmt.Println("from comma ok ", i)
			}
		}
	}
}

func send(e, o, q chan<- int) {
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
