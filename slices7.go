package main

import "fmt"

func main() {
	truck := []string{"Electric", "14000lbs of tow", "30x Cold rolled steel", "Made in America"}
	suv := []string{"Electric", "Seats 5", "Safest car in class", "Made in America"}
	fmt.Println(truck)
	fmt.Println(suv)

	myMap := [][]string{truck, suv}
	fmt.Println(myMap)
}
