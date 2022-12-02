package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(bytes)))
	score := 0
	for scanner.Scan() {
		line := scanner.Text()
		score += playRockPaperScissors(line)
	}

	fmt.Println("Total score is", score)
}

func playRockPaperScissors(line string) int {
	tokens := strings.Fields(line)
	opToken := tokens[0]
	myToken, roundScore := mapMyMoveToToken(tokens[1], opToken)
	tokenScore := mapTokenToScore(myToken)
	return (roundScore + tokenScore)
}

func mapMyMoveToToken(roundResult string, opToken string) (string, int) {
	myToken := "A"
	roundScore := 0
	switch {
	case roundResult == "X": // I need to lose
		myToken = string(opToken[0] - 1)
		roundScore = 0
	case roundResult == "Y": // I need to draw
		myToken = opToken
		roundScore = 3
	case roundResult == "Z": // I need to win
		myToken = string(opToken[0] + 1)
		roundScore = 6
	}
	// Make sure we didn't go out of bounds
	if myToken < "A" {
		myToken = "C"
	} else if myToken > "C" {
		myToken = "A"
	}
	return myToken, roundScore
}

func mapTokenToScore(token string) int {
	return int(token[0]-"A"[0]) + 1
}
