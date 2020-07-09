package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanIn(c1, c2)
}
