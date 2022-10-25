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
		"E": 5,
		"F": 6,
		"G": 7,
		"H": 8,
		"I": 9,
	}

	fmt.Println(testMap["E"])
	testMap["E"] = 420
	fmt.Println(testMap["E"])
	testMap["J"] = 420902000
	fmt.Println(testMap)
	delete(testMap, "G")
	fmt.Println(testMap)

}
