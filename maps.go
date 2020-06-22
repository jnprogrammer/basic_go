//maps are key value store, that allow for fast lookup
//maps are unordered lists

package main

import "fmt"

func main() {
	m := map[string]int{
		"CyberTruck": 804,
		"Model Y":    508,
	}

	fmt.Println(m)
	fmt.Println(m["CyberTruck"])
	//if you request a key that isn't in the map, it will return 0

	value, ok := m["Ford F-250"]
	//o, oka := m["Trap"]
	fmt.Println(value)
	fmt.Println(ok)
	//fmt.Println(o)
	//fmt.Println(oka)

	// How to check if a key is in a map in idiomatic go

	if value, ok := m["CyberTruck"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("Key not Found")
	}

}
