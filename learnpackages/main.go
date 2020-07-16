package main

import (
	"fmt"
	"learnpackages/simpleinterest"
)

func main() {

	fmt.Println(" Simple interest calculator")
	fmt.Println(" Retyped from Blog")
	fmt.Println("https://golangbot.com/go-packages/")
	fmt.Println("To re-enforce learning go packages")

	p := 5000.0
	r := 10.0
	t := 15.0

	si := simpleinterest.Calculate(p, r, t)
	fmt.Println("Simple interest is", si)
}
