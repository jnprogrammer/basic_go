package main

import (
	"fmt"
)

type shape interface {
	area()
}

type circle struct {
	r  float64
	pi float64
}

type square struct {
	l int
	w int
}

func main() {
	circle64 := circle{
		r:  42.2,
		pi: 3.14,
	}

	square94 := square{
		l: 43,
		w: 89,
	}

	shape.area(&circle64)
	shape.area(&square94)
}

func (c *circle) area() {
	a := c.pi * c.r
	fmt.Println("The Area of this circule is : ", a)
}

func (s *square) area() {
	a := s.l * s.w
	fmt.Print("The Area of the Square is: ", a)
}
