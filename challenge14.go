// print out modulus found for each number between 10 and 100 when divided by 4
package main

import "fmt"

func main() {
	number := 10
	for number <= 100 {
		mod := number % 4
		fmt.Println(mod)
		number++
	}
}
