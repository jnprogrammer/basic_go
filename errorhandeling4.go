package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("Names.txt")
	if err != nil {
		fmt.Println(err)
		return
	} //this is how you handel errors, always handel them 99.9% of the time
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
