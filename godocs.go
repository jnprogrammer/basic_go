//Package mymath provides the best math in the country
package mymath

// Sum adds an unlimited number of values of type int
func Sum(x . . .int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return  sum
}