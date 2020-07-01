package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	pw := `4fgdvFjj**`
	bs, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)
}
