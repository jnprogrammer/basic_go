package main

import (
	"fmt"
	"sort"
)

type Carrot struct {
	Lenghcm float32
	Weightg float32
	Strain  string
	Color   string
}

type ByCarrot []Carrot

func (c ByCarrot) Len() int           { return len(c) }
func (c ByCarrot) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ByCarrot) Less(i, j int) bool { return c[i].Strain < c[j].Strain }

type ByWeight []Carrot

func (w ByWeight) Len() int           { return len(w) }
func (w ByWeight) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }
func (w ByWeight) Less(i, j int) bool { return w[i].Weightg > w[j].Weightg }

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
	//fmt.Println(patch)

	sort.Sort(ByCarrot(patch))
	fmt.Println(patch)

	sort.Sort(ByWeight(patch))
	fmt.Println(patch)

}

func (c Carrot) String() string {
	return fmt.Sprintf("Carrot: %s  color: %s  weight: %vg length: %vkm \n", c.Strain, c.Color, c.Weightg, c.Lenghcm)
}
