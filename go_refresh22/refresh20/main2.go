package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	func() {
		defer waitGrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Hello Go for the 3220 time\nI KEEP FIGHTING TO CHANGE AND I KEEP FIGHTING")
	}()

	func() {
		defer waitGrp.Done()
		fmt.Println("Cats are not food")
	}()
	waitGrp.Wait()
}
