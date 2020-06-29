package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "Hey video games !")
	io.WriteString(os.Stdout, "HEy again!")
}

// this code above is used with the Go lang documentation opened
// you'll be able to find what belongs to the writer interface such as
// anything of Type File since file has a method that puts it in
// the writer interface.
