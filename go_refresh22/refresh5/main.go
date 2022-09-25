package main

import (
	"fmt"
)

func main() {
	name := "Josh"
	course := "Getting on With Go"

	fmt.Println("\nHi", name, "Your current course is: ", course)
	updateCourse(&course)

	fmt.Println("your're currently watching", course)

}

func updateCourse(course *string) string {
	*course = "Getting on with Docker"
	fmt.Println("Update course to", *course)
	return *course
}
