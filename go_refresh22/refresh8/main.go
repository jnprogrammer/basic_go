package main

import (
	"fmt"
	"strings"
)

func main() {
	author := "Ada Lovelace"
	course := "Getting a Node running"

	fmt.Println(converter(author, course))
}

func converter(author string, course string) (a, b string) {
	author = strings.ToUpper(author)
	course = strings.Title(course)

	return author, course
}
