package main

import "fmt"

func main() {
	//for i := 0; i < 420; i++ {
	//	if i % 2 != 0 {
	//		fmt.Println(i)
	//	}
	//}

	//case statements in go don't have fallthrough in switch statements less specified

	switch {
	case false:
		fmt.Print("Won't print")
	case (2 == 4):
		fmt.Println("also won't print")
	case (5 > 4):
		fmt.Println("PRints")
		fallthrough
	case false:
		fmt.Print("Prints anyway because of fallthrough in previous case\n")
		//fallthrough
	case true:
		fmt.Print("True so it prints only if there is a fallthrough in the previous step")

	}
}
