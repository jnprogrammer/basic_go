package main

import (
	"encoding/json"
	"fmt"
)

type rock struct {
	Kind  string
	Power int
}

func main() {
	u1 := rock{
		Kind:  "Fire Stone",
		Power: 342,
	}

	u2 := rock{
		Kind:  "Water Stone",
		Power: 27,
	}

	u3 := rock{
		Kind:  "Electric Stone",
		Power: 420,
	}

	rocks := []rock{u1, u2, u3}

	fmt.Println(rocks)

	// your code goes here

	bs, err := json.Marshal(rocks)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

}
