// create a slice of slice of string ([][]string) and print out the slice of a slice
package main

import "fmt"

func main() {
	x := []string{"Cats", "Dogs", "Birds"}
	y := []string{"Mice", "Cats", "Worms"}
	m := [][]string{x, y}

	fmt.Println(m)

	for i, record := range m {
		fmt.Println("Record: \n", i)
		for j, v := range record {
			fmt.Printf("\t Index position: %v \t value: %v\n", j, v)
		}
	}
}
