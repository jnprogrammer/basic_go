package main

import (
	"fmt"
	"sync"
)

var waitgroups sync.WaitGroup

func main() {

	waitgroups.Add(2)
	go func() {
		fmt.Println("Hello from this func")
		waitgroups.Done()
	}()
	go func() {
		fmt.Println("Hey from the other func!")
		waitgroups.Done()
	}()

	waitgroups.Wait()

}
