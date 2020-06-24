package main

import "fmt"

type person struct {
	name  string
	level int
	money float32
}

func main() {
	tim := person{
		name:  "Tim",
		level: 3,
		money: 435000.43,
	}

	fmt.Println(tim)
}
