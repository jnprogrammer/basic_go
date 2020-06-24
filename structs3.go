package main

import "fmt"

func main() {
	// This is an anonymous struct

	dude := struct {
		name  string
		level int
		car   string
	}{
		name:  "Tom",
		level: 5,
		car:   "Model S",
	}

	fmt.Println(dude)
}

//In Go, we create VALUES of a certain TYPE that are stored in VARIABLES
//and those VARIABLES have IDENTIFIERS
