// make a program that uses a switch expression specified as a variable of TYPE string with IDENTIFIER "favSport
package main

import "fmt"

var favSport string

func main() {

	favSport = "StarCraft 2"

	switch favSport {
	case "Basket Ball":
		fmt.Println(favSport)
	case "StarCraft 2":
		fmt.Println(favSport)
	}
}
