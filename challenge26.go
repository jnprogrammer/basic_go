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

	m := map[string]person{
		tom.last:  tom,
		bank.last: bank,
	}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
		fmt.Println(tom.icecream, bank.icecream)
	}
}
