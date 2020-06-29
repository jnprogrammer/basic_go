package main

import (
	"encoding/json"
	"fmt"
)

type ecar struct {
	Battery  int
	Model    string
	Topspeed int
}

func main() {
	car := ecar{
		Battery:  400,
		Model:    "Super C",
		Topspeed: 170,
	}

	cybertruck2 := ecar{
		Battery:  433,
		Model:    "Cyber Truck",
		Topspeed: 230,
	}

	cars := []ecar{car, cybertruck2}

	fmt.Println(cars)

	bs, err := json.Marshal(cars)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
