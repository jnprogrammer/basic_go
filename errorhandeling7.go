package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("nofilehere.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

func foo() {
	fmt.Println("when os.Exit() is called, deferred functionns don't run")
}
