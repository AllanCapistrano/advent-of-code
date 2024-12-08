package main

import (
	"fmt"

	firstchallenge "gihub.com/allancapistrano/advent-of-code/challenges/1"
	secondchallenge "gihub.com/allancapistrano/advent-of-code/challenges/2"
)

func main() {
	firstChallengeResult := firstchallenge.FirstChallenge()
	secondChallengeResult := secondchallenge.SecondChallenge()

	fmt.Printf("First challenge answer: %.0f\n", firstChallengeResult)
	fmt.Printf("Second challenge answer: %.d\n", secondChallengeResult)
}
