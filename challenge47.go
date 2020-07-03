package main

import (
	"fmt"
	"sync"
)

var waitgroups sync.WaitGroup

func main() {

	waitgroups.Add(2)
	go foo()
	go bar()

	waitgroups.Wait()

}

func foo() {
	for i := 0; i < 4; i++ {
		fmt.Println("cat: ", i)
	}
	fmt.Println("foo done")
	waitgroups.Done()
}

func bar() {
	for i := 0; i < 4; i++ {
		fmt.Println("dog: ", i)
	}
	fmt.Println("Bar done")
	waitgroups.Done()
}
