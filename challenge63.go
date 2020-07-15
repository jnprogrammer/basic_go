package main

import "fmt"

type customerr struct {
	errInfo string
}

func (ce customerr) Error() string {
	return fmt.Sprintf("Here is go using interfaces like a gopher!: %v", ce.errInfo)
}

// any value of type customerr will implement
func main() {
	fmt.Println("Hello, dudes")
	cust := customerr{
		errInfo: "****CUSTOM ERROR****",
	}
	foo(cust)
}

func foo(e error) {
	fmt.Println("this is the eroar", e, "\n", e.(customerr).errInfo)
}
