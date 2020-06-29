package main

import "fmt"

func main() {
	func(x string) {
		fmt.Println("This is the anon func right!  ", x)
	}("Darn Right !")
}

//build and use an anonymous func
