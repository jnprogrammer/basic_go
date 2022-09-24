package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var (
	name, course = "Josh", "Master of This.exe"
	module       = "4"
	clip         = 710 // just a line showing an example.
)

func main() {
	fmt.Println("Name and course are set to ", name, " and ", course, ".")
	fmt.Println("Module and clip are set to ", module, "and", clip, ".")
	fmt.Println("Name is of type", reflect.TypeOf(name))
	fmt.Println("Module is of type", reflect.TypeOf(module))

	iModule, err := strconv.Atoi(module)
	if err == nil {
		total := iModule + clip
		fmt.Println("Module plus clip equals", total)
	}
}
