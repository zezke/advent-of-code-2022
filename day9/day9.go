package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
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

/**
 * Coordinates implementation
 */
type Coordinates struct {
	x int
	y int
}

func (c *Coordinates) String() string {
	return fmt.Sprintf("(%d, %d)", c.x, c.y)
}

func (c *Coordinates) Move(action string, amount int) {
	switch action {
	case "U":
		c.Up(amount)
	case "D":
		c.Down(amount)
	case "L":
		c.Left(amount)
	default:
		c.Right(amount)
	}
}

func (c *Coordinates) Up(amount int) {
	c.y += amount
}

func (c *Coordinates) Down(amount int) {
	c.y -= amount
}

func (c *Coordinates) Left(amount int) {
	c.x -= amount
}

func (c *Coordinates) Right(amount int) {
	c.x += amount
}

/**
 * Actual solution
 */
const NR_OF_KNOTS = 10

var knotPositions []Coordinates
var tailKnotIndex int
var tailPositions *set

func main() {
	bytes, err := os.ReadFile("input")
	if err != nil {
		fmt.Println(err)
	}

	knotPositions = make([]Coordinates, NR_OF_KNOTS)
	for i := 0; i < NR_OF_KNOTS; i++ {
		knotPositions[i] = Coordinates{x: 0, y: 0}
	}
	tailKnotIndex = NR_OF_KNOTS - 1
	tailPositions = NewSet()
	tailPositions.Add(knotPositions[tailKnotIndex].String())

	// fmt.Println("== Initial State ==")
	// fmt.Println("H:", headPosition.String())
	// fmt.Println("T:", tailPosition.String())
	// fmt.Println()

	scanner := bufio.NewScanner(strings.NewReader(string(bytes)))
	for scanner.Scan() {

		line := scanner.Text()
		lineSplit := strings.Fields(line)
		move := lineSplit[0]
		amount, _ := strconv.Atoi(lineSplit[1])

		fmt.Println("== ", line, " ==")
		for i := 0; i < amount; i++ {
			DoMove(move)
		}
		fmt.Println()
	}

	fmt.Println("Unique tail positions:", tailPositions.Length())
}

func DoMove(move string) {
	knotPositions[0].Move(move, 1)

	for i := 0; i < tailKnotIndex; i++ {
		// fmt.Println("Processing move", move, "between knot", i, "and knot", i+1)
		doMoveFor(&knotPositions[i], &knotPositions[i+1])
		// fmt.Println("Result:", knotPositions[i].String(), ",", knotPositions[i+1].String())
	}

	// Store tail position
	tailPositions.Add(knotPositions[tailKnotIndex].String())
}

func doMoveFor(headPosition *Coordinates, tailPosition *Coordinates) {
	// If the  tail is ever two steps directly U,D,L or R from tail, tail must move one step in that direction
	// Check left or right
	if headPosition.y == tailPosition.y || headPosition.x == tailPosition.x {
		if headPosition.y == tailPosition.y {
			if headPosition.x == tailPosition.x+2 { // Move tail to the right
				tailPosition.Right(1)
			} else if headPosition.x == tailPosition.x-2 { // Move tail to the left
				tailPosition.Left(1)
			}
		}
		// Check up or down
		if headPosition.x == tailPosition.x {
			if headPosition.y == tailPosition.y+2 { // Move tail up
				tailPosition.Up(1)
			} else if headPosition.y == tailPosition.y-2 { // Move tail down
				tailPosition.Down(1)
			}

		}
	}

	// If the head and tail aren't touching and aren't in the same row or column, the tail always moves one step diagonally to keep up
	if headPosition.x != tailPosition.x && headPosition.y != tailPosition.y { // They should not be in the same row or column
		if math.Abs(float64(headPosition.x)-float64(tailPosition.x)) > 1 || math.Abs(float64(headPosition.y)-float64(tailPosition.y)) > 1 { // They should not be touching (could be in one if but that's ugly)
			// Check if we need to move left or right diagonally
			if tailPosition.x < headPosition.x { // Should move to the right
				tailPosition.Right(1)
			} else if tailPosition.x > headPosition.x { // Should move to the left
				tailPosition.Left(1)
			}
			// Check if we need to move up or down diagonally
			if tailPosition.y < headPosition.y {
				tailPosition.Up(1)
			} else if tailPosition.y > headPosition.y {
				tailPosition.Down(1)
			}
		}
	}
}
