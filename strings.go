package main

import (
	"fmt"
)

func main() {
	rawstring := `"Hello, strings in\n \nt \nt \n \n \n n\   /\.sfg.\smmks go"` //backtics make a raw string literal
	s := "Hello again String world"

	fmt.Printf("%T", s)
	fmt.Println(s)
	fmt.Printf("%T", rawstring)
	fmt.Println(rawstring)

	bs := []byte(rawstring) // takes a slice of bytes from string and coverts to UTF number
	fmt.Println(bs)
	fmt.Printf("%T", bs)

	//Prints the utf8 number for each char in string

	for i := 0; i < len(rawstring); i++ {
		fmt.Printf("%#U\n", rawstring[i])
	}
}
