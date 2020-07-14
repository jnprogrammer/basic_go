package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("not.txt")
	if err != nil {
		fmt.Println("err happened", err)
		//fmt.Println("err happened !?!", err)
		//log.Fatalln(err) // calls println then calls os.Exit(1)
		//panic(err)
	}
}
