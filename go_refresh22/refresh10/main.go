package main

import (
	"fmt"
)

func main() {

	if dLenMin, cLenMin := 275, 30; dLenMin > cLenMin {
		fmt.Println("D lenth is longer")
	} else if dLenMin == cLenMin {
		fmt.Println("Both are the same")
	} else {
		fmt.Println("C must be longer")
	}
}
