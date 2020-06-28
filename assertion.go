package main

import "fmt"

type electricvechical struct {
	motors              int
	batterycap          int
	manufacturer        string
	model               string
	chargenetworkaccess bool
	drivetrain          string
	topspeedkm          int
}

type electrictruck struct {
	electricvechical
	towcap          int
	groundclearance int
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
		},
		towcap:          14000,
		groundclearance: 16,
	}

	evechical := electricvechical{
		motors:              2,
		batterycap:          1020,
		manufacturer:        "Tesla",
		model:               "Model Y",
		chargenetworkaccess: true,
		drivetrain:          "Duel Motor All-Wheel Drive",
		topspeedkm:          150,
	}

	fmt.Println(ctruck.model)
	fmt.Println(evechical.model)
	ctruck.spin()
	evechical.spin()

	bar(ctruck)
	bar(evechical)
}

type trucks interface {
	spin()
}

func bar(t trucks) {
	switch t.(type) {
	case electrictruck:
		fmt.Println("I'm a Model Y made by", t.(electrictruck).manufacturer) // here I'm asserting this VALUE will be TYPE electrictruck
	case electricvechical:
		fmt.Println("I'm a CyberTryck, made by ", t.(electricvechical).manufacturer)
	}
}

func (e electrictruck) spin() {
	fmt.Println("We are spinning in a method", e.towcap)
}

func (ce electricvechical) spin() {
	fmt.Println("We are spinning in a method", ce.model)

}
