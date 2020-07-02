package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("OS:\t ", runtime.GOOS)
	fmt.Println("ARCH:\t ", runtime.GOARCH)
	fmt.Println("CPUs:\t ", runtime.NumCPU())
	fmt.Println("Go routines:\t ", runtime.NumGoroutine())
	fmt.Println("************************************")

	wg.Add(1)
	go foo()
	bar()

	fmt.Println("************************************")
	fmt.Println("CPUs:\t ", runtime.NumCPU())
	fmt.Println("Go routines:\t ", runtime.NumGoroutine())
	fmt.Println("************************************")
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("cat: ", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("dog: ", i)
	}
}
