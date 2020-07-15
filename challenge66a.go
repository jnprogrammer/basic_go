package main

import "fmt"

type dog struct{
	name string
	age int
}

func main() {
	k9 := dog{
		name: "Dogguo",
		age:3,
	}

	fmt.Println(k9)
}