package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(bytes)))
	elfCalories := make([]int, 0)
	currentCalories := 0
	for scanner.Scan() {
		line := scanner.Text()
		// Check if we have reached an empty line (file contains empty line at end as well)
		if len(strings.TrimSpace(line)) == 0 {
			elfCalories = append(elfCalories, currentCalories)
			currentCalories = 0
			continue
		}
		// Otherwise we sum up the calories
		calories, _ := strconv.Atoi(line)
		currentCalories += calories
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elfCalories)))
	fmt.Println(elfCalories[0:3])

	sum := elfCalories[0] + elfCalories[1] + elfCalories[2]
	fmt.Println("Sum is", sum)
}
