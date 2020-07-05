package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var mu sync.Mutex
	var wg sync.WaitGroup
	var counter int
	const goes = 100

	wg.Add(goes)
	for i := 0; i < goes; i++ {
		go func() {
			mu.Lock()
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			//)
			wg.Done()

		}()
	}
	wg.Wait()
	fmt.Println("count:", counter)

	fmt.Println("Those Go Routines ?", runtime.NumGoroutine())

}
