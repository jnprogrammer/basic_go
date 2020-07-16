//Package mymath provides the best math in the country
package mymath

// Sum adds an unlimited number of values of type int
func Sum(x int . . . int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return  sum
}

/*
So the coolest thing about go when it comes to documentation
is that when you create comments and upload your code to github, you can
have that code searchable in godocs, and go docs will display them formatted


Remember when making documentation you start with the name of what your
documenting first and write in complete sentences.
 */