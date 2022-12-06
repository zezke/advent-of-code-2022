package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
 * Set implementation
 */
var exists = struct{}{}

type set struct {
	m map[string]struct{}
}

func NewSet() *set {
	s := &set{}
	s.m = make(map[string]struct{})
	return s
}

func (s *set) Add(value string) {
	s.m[value] = exists
}

func (s *set) Remove(value string) {
	delete(s.m, value)
}

func (s *set) Contains(value string) bool {
	_, c := s.m[value]
	return c
}

func (s *set) Length() int {
	return len(s.m)
}

func (s *set) Print() string {
	var sb strings.Builder
	sb.WriteString("Set: {")
	for str := range s.m {
		sb.WriteString(str)
		sb.WriteString(", ")
	}
	sb.WriteString("}")
	return sb.String()
}

/**
 * Actual solution
 */
const DISTINCT_CHAR_LENGTH = 14

func main() {
	bytes, err := os.ReadFile("input")
	if err != nil {
		fmt.Println(err)
	}

	// var cranes []Stack
	scanner := bufio.NewScanner(strings.NewReader(string(bytes)))
	for scanner.Scan() {
		line := scanner.Text()
		firstMarkerPos := FindFirstMarkerInLine(line)
		fmt.Println("First marker pos:", firstMarkerPos)
	}
}

func FindFirstMarkerInLine(line string) int {

	// fmt.Println("FindFirstMarkerInLine:", line)
	for i := 0; i < len(line); i++ {
		// Build set with this pos char + DISTINCT_CHAR_LENGTH preceeding
		set := NewSet()
		startIndex := i - DISTINCT_CHAR_LENGTH + 1
		if startIndex < 0 {
			startIndex = 0
		}
		// fmt.Println("    At pos", i, "line is", line[startIndex:i+1])
		for startIndex <= i {
			char := string(line[startIndex])
			set.Add(char)
			startIndex++
		}
		char := string(line[i])
		set.Add(char)
		// fmt.Println("    At pos", i, ", the set is", set.Print())

		// Check if we have DISTINCT_CHAR_LENGTH unique values
		if set.Length() == DISTINCT_CHAR_LENGTH {
			return i + 1
		}
	}
	return -1
}
