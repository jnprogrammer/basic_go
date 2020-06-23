//take a slice of a slice
package main

import (
	"fmt"
)

func main() {
	comp := []int{45, 53, 2, 420, 4522}

	fmt.Println(comp[:4])
	fmt.Println(comp[2:4])
	fmt.Println(comp[:1])
	fmt.Println(comp[4:])

}
