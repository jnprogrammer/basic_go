package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Carrot struct {
	Lenghcm float32
	Weightg float32
	Strain  string
	Color   string
}

func main() {
	redcarrots := Carrot{
		Strain:  "Firebrand",
		Color:   "Neon Red",
		Weightg: 4.32,
		Lenghcm: 2.23,
	}

	purplecarrots := Carrot{
		Strain:  "soothingpurps",
		Color:   "Dark Purple",
		Weightg: 6.12,
		Lenghcm: 3.40,
	}

	greencarrots := Carrot{
		Strain:  "druidsbroom",
		Color:   "Bright Green",
		Weightg: 3.86,
		Lenghcm: 5.237,
	}

	pinkcarrots := Carrot{
		Strain:  "candyclaw",
		Color:   "Neon Pink",
		Weightg: 8.964,
		Lenghcm: 5.22,
	}

	patch := []Carrot{redcarrots, purplecarrots, greencarrots, pinkcarrots}
	fmt.Println(patch)

	err := json.NewEncoder(os.Stdout).Encode(patch)
	if err != nil {
		fmt.Println("We don't have any carrots ?")
	}

}
