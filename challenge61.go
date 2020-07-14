package main

import (
	"encoding/json"
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
		log.Fatalln("Error in marshaling your string", err)
	}
	fmt.Println(string(bs))

}
