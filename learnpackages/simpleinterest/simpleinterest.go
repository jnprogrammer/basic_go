package simpleinterest

import "fmt"

func init() {
	fmt.Println("Running Simple interest package")
}

//Calculate calculates and returns the simple interest for principal p, rate of interest r for time duration t years
func Calculate(p float64, r float64, t float64) float64 {
	interest := p * (r / 100) * t
	return interest
}
