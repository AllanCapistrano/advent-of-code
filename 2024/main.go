package main

import (
	"fmt"

	challenges "gihub.com/allancapistrano/advent-of-code/challenges/1"
	secondchallenges "gihub.com/allancapistrano/advent-of-code/challenges/2"
)

func main() {
	firstChallengeResult := challenges.FirstChallenge()
	secondChallengeResult := secondchallenges.SecondChallenge()

	fmt.Printf("First challenge answer: %.0f\n", firstChallengeResult)
	fmt.Printf("Second challenge answer: %.d\n", secondChallengeResult)
}
