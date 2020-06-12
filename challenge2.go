package main

import (
	"fmt"
)

var x int = 420
var y string = "sudo apt-get thc"
var z bool = true

func main() {

	// adds all objects to a single string.
	singlestr := fmt.Sprintf("%v\n%v\n%v", x, y, z)
	fmt.Println(singlestr)
}
