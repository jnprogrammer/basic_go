package main

import "fmt"

type dude struct {
	first string
	last  string
}

type alladude struct {
	dude
	acooldude bool
}

func main() {
	Kel := alladude{
		dude: dude{
			"Kel",
			"Orange LOVES Soda",
		},
		acooldude: true,
	}

	Ken := alladude{
		dude: dude{
			"Ken",
			"Red LOVES Soup",
		},
		acooldude: true,
	}

	fmt.Println(Kel)
	Kel.speak()
	Ken.speak()
}

// Any VALUE of TYPE alladude now has access to the speak() function
func (s alladude) speak() {
	fmt.Println("I'm a Dude! ", s.first, s.last)
}

// func (r receiver) identifier(parameters) (return(s)) { code }
