package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	var wg sync.WaitGroup
	var counter int64
	const goes = 100

	wg.Add(goes)
	for i := 0; i < goes; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()

		}()
	}
	wg.Wait()
	fmt.Println("count:", counter)

	fmt.Println("Those Go Routines ?", runtime.NumGoroutine())

}
