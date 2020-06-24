package main

import "fmt"

type person struct {
	name  string
	level int
	money float32
}

type human struct {
	person
	soul   bool
	age    int
	cares  bool
	nature string
}

func main() {
	s := human{
		person: person{
			name:  "Tim",
			level: 100,
			money: 4352.42,
		},
		soul:   true,
		age:    24,
		cares:  true,
		nature: "air",
	}

	fmt.Println(s)
	fmt.Println(s.name, s.age, s.money, s.name, s.level)

}
