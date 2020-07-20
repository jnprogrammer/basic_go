package challenge67

import (
	"fmt"
	"go/challenge67/ben"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  ben.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(ben.YearsTwo(20))
}
