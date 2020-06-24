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
