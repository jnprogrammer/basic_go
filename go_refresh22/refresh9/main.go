package main

import (
	"fmt"
)

func main() {
	bestFinish := champFinishes(1, 12, 5, 6, 4, 3, 3)

	fmt.Println(bestFinish)
}

func champFinishes(finishes ...int) int {
	best := finishes[0]
	for _, i := range finishes {
		if i < best {
			best = i
		}
	}
	return best
}
