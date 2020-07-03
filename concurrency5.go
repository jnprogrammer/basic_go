package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GOROUTINES:", runtime.NumGoroutine())

	var waitgroups sync.WaitGroup

	var counter int64

	const gos = 100
	waitgroups.Add(gos)

	for i := 0; i < gos; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			waitgroups.Done()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())

	}

	waitgroups.Wait()
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("count:", counter)
	fmt.Println("Hello, concurrent Go !")
}

//Learning about Atomic package
