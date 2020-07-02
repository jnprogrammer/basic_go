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

	bs, err := json.Marshal(rocks)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

	//unmarshal the data from the previous challenge
	err = json.Unmarshal(bs, &rocks)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("All the data: ", rocks)

	for i, v := range rocks {
		fmt.Println("\n Rocks: ", i)
		fmt.Println(v.Kind, v.Power)
	}
}
