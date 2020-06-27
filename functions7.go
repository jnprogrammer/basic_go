package main

import (
	"fmt"
)

type electricvechical struct {
	motors              int
	batterycap          int
	manufacturer        string
	model               string
	chargenetworkaccess bool
	drivetrain          string
	topspeedkm          int
	autopilot           bool
}

type electrictruck struct {
	electricvechical
	towcap          int
	groundclearance int
}

type ecars interface {
	honk()
}

func main() {
	ctruck := electrictruck{
		electricvechical: electricvechical{
			motors:              3,
			batterycap:          806,
			manufacturer:        "Tesla",
			model:               "CyberTruck",
			chargenetworkaccess: true,
			drivetrain:          "Tri Motor All-Wheel Drive",
			topspeedkm:          120,
			autopilot:           true,
		},
		towcap:          14000,
		groundclearance: 16,
	}
	atruck := electrictruck{
		electricvechical: electricvechical{
			motors:              1,
			batterycap:          450,
			manufacturer:        "Tesla",
			model:               "CyberTruck",
			chargenetworkaccess: true,
			drivetrain:          "Single Motor Rear-Wheel Drive",
			topspeedkm:          120,
			autopilot:           true,
		},
		towcap:          7000,
		groundclearance: 16,
	}
	ctruck.honk()

	c := ecars(ctruck)
	a := ecars(atruck)

	fmt.Println(a, c)
}

func (s electrictruck) honk() {
	fmt.Println("This is where I'm from and what made me: ", s.manufacturer, s.model, s.drivetrain)
}

// A VALUE can be of more than one TYPE

func show(e ecars) {
	fmt.Println("I was passed into show", e)
}
