package main

import "fmt"

func main() {
	m := map[string]int{
		"CyberTruck": 804,
		"Model Y":    508,
	}
	fmt.Println(m)
	// How to add to a map
	m["Model X"] = 420

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}

	xi := []int{4, 5, 6, 7, 8, 9, 0}

	for i, v := range xi {
		fmt.Println(i, v)
	}

	// composite literals are used to build maps.

	// to delete a key from a map use delete(<map name>, "key")

	//delete(m, "Model X")
	//
	//delete(m, "Captain Planet")
	//if you try to delete a key that isn't there the code will work, however it will not alert you that the key didn't exist
	// use the Golang Comma Ok idiom to test for this

	if v, ok := m["Model X"]; ok {
		fmt.Println("value: ", v)
		delete(m, "Model X")
	}
	for k, v := range m {
		fmt.Println(k, v)
	}

}
