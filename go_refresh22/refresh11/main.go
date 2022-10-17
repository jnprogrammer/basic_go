package main

func main() {
	type courseMeta struct {
		auth  string
		level string
		rate  float64
	}

	getStartedWinning := new(courseMeta)

	getStartedWinning.auth = "me"

	BeWinning := courseMeta{
		auth:  "Tim",
		level: "Intermed",
		rate:  45,
	}

	println(getStartedWinning.auth, BeWinning.auth)

}
