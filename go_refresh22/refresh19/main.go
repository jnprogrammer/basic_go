package main

import "fmt"

func main() {
	leagueTitles := make(map[string]int)
	leagueTitles["SunderLAND"] = 6
	leagueTitles["Newcastle"] = 4

	recentWins := map[string]int{
		"sunderland": 4,
		"Timmy":      45,
	}

	fmt.Printf("Titles: %v\nRecent wins: %v\n", leagueTitles, recentWins)
}
