package main

import (
	"benchmark2/bench"
	"fmt"
	"strings"
)

const s = "cats"

func main() {
	strslice := strings.Split(s, " ")

	for _, v := range strslice {
		fmt.Println(v)
	}

	fmt.Printf("\n%s\n", bench.Cat(strslice))
	fmt.Printf("\n%s\n\n", bench.Join(strslice))
}
