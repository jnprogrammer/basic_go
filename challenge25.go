// create a struct and store some data
// create two values of that type

package main

import (
	"fmt"
)

type person struct {
	first    string
	last     string
	icecream []string
}

func main() {

	tom := person{
		first:    "Tom",
		last:     "Soyer",
		icecream: []string{"Cookie Dough", "CAts"},
	}

	bank := person{
		first:    "WCA",
		last:     "CEA",
		icecream: []string{"Rocky Roat", "berry"},
	}

	fmt.Println(tom.first, tom.last, tom.icecream)
	fmt.Println(bank.first, bank.last)
	for i, v := range bank.icecream {
		fmt.Println(i, v)
	}
}
