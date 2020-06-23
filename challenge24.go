// create a map with key TYPE string, store three records in the map, print out all values with their index too.

package main

import "fmt"

func main() {
	m := map[string]int{
		"Water": 1,
		"Earth": 2,
		"Fire":  3,
		"Air":   4,
	}
	fmt.Println("****************")

	for key, val := range m {
		fmt.Println(key, val)
	}
	fmt.Println("****************")
	// adding a record

	m["Avatar"] = 5
	m["Zuko"] = 420

	for key, val := range m {
		fmt.Println(key, val)
	}
	fmt.Println("****************")

	//if v, ok := m["Zuko"]; ok {
	//	delete(m, "Zuko")
	//}
	fmt.Println("Delete Zuko")
	delete(m, "Zuko")
	fmt.Println("****************")

	for key, val := range m {
		fmt.Println(key, val)
	}

}
