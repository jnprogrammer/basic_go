// make a typed and untyped const

package main

import "fmt"

const range_in_miles int = 500

const (
	cars  = "Model Y"
	truck = "CyberTruck"
)

func main() {

	fmt.Printf("This is a electric car: %q\n", cars)
	fmt.Printf("This is a electric truck: %q and its range %d", truck, range_in_miles)
}
