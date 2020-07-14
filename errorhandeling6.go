package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.text")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("nofile.txt")
	if err != nil {
		log.Println("Err happened", err)
	}
	defer f2.Close()

	fmt.Println("check the log.txt for error information")
}
