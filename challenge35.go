//assign a func to a variable

package main

import "fmt"

func main() {
	afunction := func() {
		fmt.Println("This was assigned to a variable")
	}

	afunction()
}
