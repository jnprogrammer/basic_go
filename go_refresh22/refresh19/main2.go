package main

import (
	"fmt"
)

func main() {

	testMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
	}

	for mapKey, mapVal := range testMap {
		fmt.Printf("Key is: %v Value is: %v\n", mapKey, mapVal)
	}
}
