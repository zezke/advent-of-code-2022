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
	sum := 0
	for scanner.Scan() {
		// Get the groups rucksack contents
		var groupRucksacks [3]string
		groupRucksacks[0] = scanner.Text()
		for i := 1; i < 3; i++ {
			scanner.Scan()
			groupRucksacks[i] = scanner.Text()
		}

		// Find the intersection between them
		firstIntersection := findIntersection([]rune(groupRucksacks[0]), []rune(groupRucksacks[1]))
		secondIntersection := findIntersection([]rune(groupRucksacks[1]), []rune(groupRucksacks[2]))
		finalIntersection := findIntersection(firstIntersection, secondIntersection)

		sum += badgeToScore(finalIntersection[0])
	}

	fmt.Println("Sum of the priorities is", sum)
}

func badgeToScore(badge rune) int {
	// fmt.Println("Determining score for badge", string(badge))
	if int(badge) >= int('a') { // a starts at 61
		return int(badge) - int('a') + 1
	}
	return int(badge) - int('A') + 27
}

func findIntersection(chars1, chars2 []rune) (inter []rune) {
	hash := make(map[rune]bool)
	for _, e := range chars1 {
		hash[e] = true
	}
	for _, e := range chars2 {
		// If elements present in the hashmap then append intersection list.
		if hash[e] {
			inter = append(inter, e)
		}
	}
	//Remove dups from slice.
	inter = removeDuplicates(inter)
	return
}

func removeDuplicates(elements []rune) (nodups []rune) {
	encountered := make(map[rune]bool)
	for _, element := range elements {
		if !encountered[element] {
			nodups = append(nodups, element)
			encountered[element] = true
		}
	}
	return
}
