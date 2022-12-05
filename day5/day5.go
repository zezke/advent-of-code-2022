package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 * Stack implementation to make my life a bit easier on day 5
 */
type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}

func (s *Stack) Peek(pos int) (string, bool) {
	if pos >= len(*s) {
		return "", false
	}
	element := (*s)[pos]
	return element, true
}

/**
 * Actual solution
 */
func main() {
	bytes, err := os.ReadFile("input")
	if err != nil {
		fmt.Println(err)
	}

	// var cranes []Stack
	scanner := bufio.NewScanner(strings.NewReader(string(bytes)))
	// Deal with the crane input
	craneData := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) <= 0 { //
			break
		}
		craneData = append(craneData, line)
	}
	craneNumbers := strings.Fields(craneData[len(craneData)-1])
	numberOfCranes, _ := strconv.Atoi(craneNumbers[len(craneNumbers)-1])
	largestStack := len(craneData) - 1
	cranes := make([]Stack, numberOfCranes, numberOfCranes)
	for i := largestStack - 1; i >= 0; i-- {
		for j := 0; j < numberOfCranes; j++ {
			idx := 1 + j*4
			payload := string(craneData[i][idx])
			if len(strings.TrimSpace(payload)) > 0 {
				cranes[j].Push(payload)
			}
		}
	}
	fmt.Println("Input cranes:")
	printCranes(cranes)
	fmt.Println("-------------")

	// Deal with the movement list
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		amount, _ := strconv.Atoi(parts[1])
		source, _ := strconv.Atoi(parts[3])
		destination, _ := strconv.Atoi(parts[5])
		move(cranes, amount, source, destination)
		printCranes(cranes)
	}
}

func move(cranes []Stack, amount int, source int, destination int) {
	fmt.Println("move", amount, "from", source, "to", destination)
	// Pt 2 solution
	payloads := make([]string, 0)
	for i := 0; i < amount; i++ {
		payload, _ := cranes[source-1].Pop()
		payloads = append(payloads, payload)
	}
	for i := len(payloads) - 1; i >= 0; i-- {
		cranes[destination-1].Push(payloads[i])
	}
}

// Pt 1
func moveSingle(cranes []Stack, source int, destination int) {
	// !! Apply -1 since index is not 0-based in input data
	payload, _ := cranes[source-1].Pop()
	cranes[destination-1].Push(payload)
}

func printCranes(cranes []Stack) {
	numberOfCranes := len(cranes)
	largestStack := 0
	for i := 0; i < numberOfCranes; i++ {
		stackSize := len(cranes[i])
		if stackSize > largestStack {
			largestStack = stackSize
		}
	}
	var sb strings.Builder
	for stackIdx := largestStack - 1; stackIdx >= 0; stackIdx-- {
		for craneIdx := 0; craneIdx < numberOfCranes; craneIdx++ {
			payload, found := cranes[craneIdx].Peek(stackIdx)
			if found {
				sb.WriteString("[")
				sb.WriteString(payload)
				sb.WriteString("] ")
			} else {
				sb.WriteString("    ")
			}
		}
		sb.WriteRune('\n')
	}
	for craneIdx := 0; craneIdx < numberOfCranes; craneIdx++ {
		sb.WriteString(" ")
		sb.WriteString(strconv.Itoa(craneIdx + 1))
		sb.WriteString("  ")
	}
	sb.WriteRune('\n')
	fmt.Print(sb.String())
}
