package main

import "fmt"

func main() {
	g := func(b []int) int {
		if len(b) == 0 {
			return 0
		}
		if len(b) == 1 {
			return b[0]
		}
		return b[0] + b[len(b)-1]
	}

	x := foo(g, []int{1, 2, 3, 4, 5, 6, 5, 7})
	fmt.Println(x)
}

func foo(f func(b []int) int, a []int) int {
	n := f(a)
	n++
	return n
}
