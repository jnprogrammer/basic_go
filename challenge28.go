//create an use an anonymous struct
package main

import "fmt"

func main() {

	modely := struct {
		battery_size int
		motors       int
		topspeed     int
	}{
		battery_size: 420,
		motors:       2,
		topspeed:     4000,
	}

	fmt.Println(modely)
	fmt.Println(modely.topspeed)

}
