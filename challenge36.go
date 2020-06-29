package main

import "fmt"

func main() {
	fk := funky()()

	fmt.Println(fk)
}

func funky() func() string {
	return func() string {
		return "This is funky returned funk"
	}
}
