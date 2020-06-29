package main

import "fmt"

func main() {
	f := func() {
		fmt.Println(":My first func expression")
	}

	f()

	g := func(s int) {
		fmt.Println(":My first number expression", s)
	}

	g(412)
}
