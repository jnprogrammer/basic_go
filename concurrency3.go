package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GOROUTINES:", runtime.NumGoroutine())

	var waitgroups sync.WaitGroup
	counter := 0

	const gos = 100
	waitgroups.Add(gos)

	for i := 0; i < gos; i++ {
		go func() {
			v := counter
			time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			waitgroups.Done()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())

	}

	waitgroups.Wait()
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("count:", counter)
	fmt.Println("Hello, concurrent Go !")
}

//This code is an example of the issue called Race Condition, you don't want these.
