package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GOROUTINES:", runtime.NumGoroutine())

	var waitgroups sync.WaitGroup
	var mu sync.Mutex //Added a mutext to lock gorutines

	counter := 0

	const gos = 100
	waitgroups.Add(gos)

	for i := 0; i < gos; i++ {
		go func() {
			mu.Lock()
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			waitgroups.Done()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())

	}

	waitgroups.Wait()
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("count:", counter)
	fmt.Println("Hello, concurrent Go !")
}

//Mutex can lock down coad to fix race conditions
