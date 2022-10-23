package main

import (
	"fmt"
	"time"
)

func main() {
	for timer := 10; timer >= 0; timer-- {
		fmt.Println(timer)
		if timer == 0 {
			fmt.Println("Boom!")
		}
		time.Sleep(1 * time.Second)
	}
}
