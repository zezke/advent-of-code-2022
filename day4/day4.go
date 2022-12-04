package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input")
	if err != nil {
		fmt.Println(err)
	}

	count := 0
	totalNumber := 0
	scanner := bufio.NewScanner(strings.NewReader(string(bytes)))
	for scanner.Scan() {
		line := scanner.Text()
		pairs := strings.Split(line, ",")

		fullyContains := processPair(pairs[0], pairs[1])
		if fullyContains {
			count += 1
		} else {
			fmt.Println(line)
		}
		totalNumber += 1
	}
	fmt.Println("Number of times that one range fully contains the other is", count, "out of a total", totalNumber)
}

func processPair(firstElf, secondElf string) bool {
	firstStart, firstEnd := stringToIndexes(firstElf)
	secondStart, secondEnd := stringToIndexes(secondElf)
	// We have an overlap if one segment starts within the other
	if firstStart >= secondStart && firstStart <= secondEnd {
		return true
	}
	if secondStart >= firstStart && secondStart <= firstEnd {
		return true
	}
	return false
}

func stringToIndexes(assignmentRange string) (int, int) {
	indexes := strings.Split(assignmentRange, "-")
	start, _ := strconv.Atoi(indexes[0])
	end, _ := strconv.Atoi(indexes[1])
	return start, end
}
