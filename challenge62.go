package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type person struct {
	First string
	Truck string
	Feats []string
}

func main() {
	p1 := person{
		First: "Tom",
		Truck: "CyberTruck",
		Feats: []string{"3 Electric Motors", "30x Steel EXOSKELETON?", "14,000+ LBS Towing cap"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))

}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("There was an error in to JSON: %v", err)

		//shows how depending on the expected type, you can chain function calls together, though the best option is the return above.
		return []byte{}, errors.New(fmt.Sprintf("There was an error in to JSON: %v", err))
	}
	return bs, nil
}
