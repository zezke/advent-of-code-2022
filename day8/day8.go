package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 * Helper struct
 */
type Tree struct {
	height    int
	isVisible bool
}

/**
 * Main logic
 */
func main() {
	bytes, err := os.ReadFile("input")
	if err != nil {
		fmt.Println(err)
	}

	// Gather the input
	forest := make([][]Tree, 0)
	scanner := bufio.NewScanner(strings.NewReader(string(bytes)))
	for scanner.Scan() {
		line := scanner.Text()
		forestLine := make([]Tree, 0)
		for _, treeHeightCh := range line {
			treeHeight, _ := strconv.Atoi(string(treeHeightCh))
			tree := Tree{height: treeHeight, isVisible: false}
			forestLine = append(forestLine, tree)
		}
		forest = append(forest, forestLine)
	}

	/** PT1 **/
	// // Look around for trees
	// // Look from the left
	// for y := 0; y < len(forest); y++ {
	// 	treeHeight := -1
	// 	for x := 0; x < len(forest[y]); x++ {
	// 		if forest[y][x].height > treeHeight {
	// 			forest[y][x].isVisible = true
	// 			treeHeight = forest[y][x].height
	// 		}
	// 	}
	// }
	// // Look from the right
	// for y := 0; y < len(forest); y++ {
	// 	treeHeight := -1
	// 	for x := len(forest[y]) - 1; x >= 0; x-- {
	// 		if forest[y][x].height > treeHeight {
	// 			forest[y][x].isVisible = true
	// 			treeHeight = forest[y][x].height
	// 		}
	// 	}
	// }
	// // Look from the top
	// for x := 0; x < len(forest[0]); x++ {
	// 	treeHeight := -1
	// 	for y := 0; y < len(forest); y++ {
	// 		if forest[y][x].height > treeHeight {
	// 			forest[y][x].isVisible = true
	// 			treeHeight = forest[y][x].height
	// 		}
	// 	}
	// }
	// // Look from the bottom
	// for x := 0; x < len(forest[0]); x++ {
	// 	treeHeight := -1
	// 	for y := len(forest) - 1; y >= 0; y-- {
	// 		if forest[y][x].height > treeHeight {
	// 			forest[y][x].isVisible = true
	// 			treeHeight = forest[y][x].height
	// 		}
	// 	}
	// }
	// PrintForest(forest)
	// fmt.Println()

	maxScore := 0
	for y := 1; y < len(forest)-1; y++ { // Don't consider the border trees since those are multiplied by zero anyway
		for x := 1; x < len(forest[y])-1; x++ {
			score := CalculateScenicValueForTree(forest, x, y)
			if score > maxScore {
				maxScore = score
			}
		}
	}
	fmt.Println("Highest scenic score for a tree is", maxScore)
}

func PrintForest(forest [][]Tree) {
	var sb strings.Builder
	for y := 0; y < len(forest); y++ {
		for x := 0; x < len(forest[y]); x++ {
			sb.WriteString(strconv.Itoa(forest[y][x].height))
			if forest[y][x].isVisible {
				sb.WriteString("X")
			} else {
				sb.WriteString("O")
			}
		}
		sb.WriteString("\n")
	}
	fmt.Print(sb.String())
}

func CountVisibleTrees(forest [][]Tree) int {
	sum := 0
	for y := 0; y < len(forest); y++ {
		for x := 0; x < len(forest[y]); x++ {
			if forest[y][x].isVisible {
				sum += 1
			}
		}
	}
	return sum
}

func CalculateScenicValueForTree(forest [][]Tree, x int, y int) int {
	// fmt.Println("Calculating scenic value for tree at (", x, ",", y, ")")
	treeHeight := forest[y][x].height
	// Look left
	leftScore := 0
	for xIt := x - 1; xIt >= 0; xIt-- {
		leftScore++
		if treeHeight <= forest[y][xIt].height {
			break
		}
	}
	// fmt.Println("Left score", leftScore)
	// Look right
	rightScore := 0
	for xIt := x + 1; xIt < len(forest[y]); xIt++ {
		rightScore++
		if treeHeight <= forest[y][xIt].height {
			break
		}
	}
	// fmt.Println("Right score", rightScore)
	// Look up
	upScore := 0
	for yIt := y - 1; yIt >= 0; yIt-- {
		upScore++
		if treeHeight <= forest[yIt][x].height {
			break
		}
	}
	// fmt.Println("Up score", upScore)
	// Look down
	downScore := 0
	for yIt := y + 1; yIt < len(forest); yIt++ {
		downScore++
		if treeHeight <= forest[yIt][x].height {
			break
		}
	}
	// fmt.Println("Down score", downScore)
	return leftScore * rightScore * upScore * downScore
}
