package main

import "fmt"

type human struct {
	name  string
	age   int
	power int
}

func main() {
	Samus := human{
		name:  "Samus",
		age:   25,
		power: 1000,
	}

	Samus.exist()

}

func (h human) exist() {
	fmt.Println("He i am, on Earth for once", h.name, h.power)
}
