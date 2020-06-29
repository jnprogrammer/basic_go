package main

import "fmt"

func main() {
	f := fib(42)
	fmt.Println(f)
}

func fib(n int) int {

	if n <= 2 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}
