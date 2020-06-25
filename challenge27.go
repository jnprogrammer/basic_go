package main

import "fmt"

type motor struct {
	motor_type string
	motors     int
}

type battery struct {
	range_km  int
	weight_kg int
}

type truck struct {
	battery
	motor
	name      string
	body_type string
}

func main() {
	cybertruck := truck{
		motor: motor{
			motor_type: "Electric",
			motors:     3,
		},
		battery: battery{
			range_km:  804,
			weight_kg: 1406,
		},
		name:      "CyberTruck",
		body_type: "30X cold-rolled stainless steel",
	}
	fmt.Println(cybertruck)
}
