package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go faninout(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("I'm gonna exit")
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func faninout(c1, c2 chan int) {
	var waitinggroup sync.WaitGroup
	for v := range c1 {
		waitinggroup.Add(1)
		go func(v2 int) {
			c2 <- timeConsumingWork(v2)
			waitinggroup.Done()
		}(v)
	}
	waitinggroup.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
