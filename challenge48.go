package main

import "fmt"

type persons struct {
	name       string
	powerlevel int
}

type humans interface {
	speak()
}

func main() {
	per =: persons{
		name: "Tom",
		powerlevel: 420
	}

	per.speak()
}

func (h persons) turnup() {
	fmt.Println("HEY I SAID MY NAME IS ")
	h.speak()
}

func (p *persons) speak(){
	fmt.Println("I'm ", p.name)
}
