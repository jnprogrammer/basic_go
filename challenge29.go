package main

import (
	"fmt"
)

func main() {
	f1 := foo(42)
	f2, f3 := boo()

	fmt.Println(f2, f3)
	fmt.Println(f1)
}

func foo(i int) int {
	return i

}

func boo() (int, string) {
	return 710, "Time to go FAST!"
}
