package main

import (
	"encoding/json"
	"fmt"
)

// Made this using https://mholt.github.io/json-to-go/
type zcar struct {
	Battery  int    `json:"Battery"`
	Model    string `json:"Model"`
	Topspeed int    `json:"Topspeed"`
}

func main() {
	j := `[{"Battery":400,"Model":"Super C","Topspeed":170},{"Battery":433,"Model":"Cyber Truck","Topspeed":230}]`
	byteslice := []byte(j)
	fmt.Printf("%T\n", j)
	fmt.Printf("%T\n", byteslice)

	//carz := []zcar{}
	var carz []zcar

	err := json.Unmarshal(byteslice, &carz)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("All of the data: ", carz)

	for i, v := range carz {
		fmt.Println("\n---- Car: ", i)
		fmt.Println(v.Model, v.Battery, v.Topspeed)
	}
}
