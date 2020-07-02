package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	pw := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	testpw := `password123`
	fmt.Println(bs)
	err = bcrypt.CompareHashAndPassword(bs, []byte(testpw))
	if err != nil{
		fmt.Println("You can';t log in!!")
	}
	fmt.Println("Your in!")
}
