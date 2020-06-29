package main

import "fmt"

type sperson struct {
	changeme int
	whoami   string
}

func main() {
	x := sperson{
		changeme: 2,
		whoami:   "Tom Riddle",
	}
	fmt.Println(x.changeme)
	fmt.Println(x)
	valchange(&x, 453)
	fmt.Println(x.changeme)
}

func valchange(y *sperson, x int) {
	y.changeme = x
}
