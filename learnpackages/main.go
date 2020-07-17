package main

import (
	"fmt"
	"learnpackages/simpleinterest"
	//_"learnpackages/simpleinterest" adding the _ before the package will silence the compiler, this os best for testing
	"log"
)

var p, r, t = -75000.0, 24.5, 10.0

//var _ = simpleinterest.Calculate() this code would allow you to run
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
	fmt.Println("Main package init")
	si := simpleinterest.Calculate(p, r, t)
	fmt.Println("Simple interest is", si)
}
