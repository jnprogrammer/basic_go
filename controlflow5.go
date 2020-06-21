package main

import "fmt"

func main() {

	//true false are predeclared constances.
	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")

	}
	if !true {
		fmt.Println("003")
	}
	if !false {
		fmt.Println("004")

	}
	if 2 == 2 {
		fmt.Println("005")
	}
	if 2 != 2 {
		fmt.Println("006")

	}
	if !(2 == 2) {
		fmt.Println("007")
	}
	if !(2 != 2) {
		fmt.Println("008")

	}

	if x := 42; x == 42 {
		fmt.Println("in Go, statements end with ; those are added by Go so doing it yourself would be redundant thus not idiomatic in Go Lang")
	}

}
