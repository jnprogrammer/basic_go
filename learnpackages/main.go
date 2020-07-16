package main

import (
	"fmt"
	"learnpackages/simpleinterest"
	"log"
)

var p, r, t = 75000.0, 24.5, 10.0

func init() {
	fmt.Println(" Simple interest calculator")
	fmt.Println(" Retyped from Blog")
	fmt.Println("https://golangbot.com/go-packages/")
	fmt.Println("To re-enforce learning go packages")

	println("main package initlize")
	if p < 0 {
		log.Fatal("Principal is less than zero")
	}
	if r < 0 {
		log.Fatal("Rate of interest is less than zero")
	}
	if t < 0 {
		log.Fatal("Duration is less than zero")
	}

}

func main() {

	si := simpleinterest.Calculate(p, r, t)
	fmt.Println("Simple interest is", si)
}
