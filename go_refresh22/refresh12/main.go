package main

import (
	"fmt"
)

func main() {

	// only exact type & matching value will activate a case
	switch "Kube" {
	case "kube":
		fmt.Println("How did you get me?")
	case "KUbe":
		fmt.Println("How did you get me?")
	case "kuBE":
		fmt.Println("How did you get me?")
	case "KubE":
		fmt.Println("How did you get me?")
	case "Kube":
		fmt.Println("we match")
	case "kubE":
		fmt.Println("How did you get me?")
	default:
		fmt.Println("Nothing?")

	}
}
