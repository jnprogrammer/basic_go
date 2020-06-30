package main

import (
	"crypto"
	"fmt"
)

func main() {
	pw := `4fgdvFjj**`
	bs := crypto.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
}
