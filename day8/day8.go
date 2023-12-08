package main

import (
	"adventofcode2023/utils"
	"fmt"
	"log"
	"strings"
)

const (
	sampleInput = "test_input.txt"
	puzzleInput = "input.txt"
)

type Day struct{}

type Path struct {
	path  string
	right string
	left  string
}

const (
	DIRECTION_LEFT  = "L"
	DIRECTION_RIGHT = "R"
	END_PATH        = "ZZZ"
)

func main() {
	day := Day{}
	solution1 := day.Part1(puzzleInput)
	fmt.Printf("Part 1 solution %d\n", solution1)

	solution2 := day.Part2(sampleInput)
	fmt.Printf("Part 2 solution %d\n", solution2)
}

func (d Day) Part1(filename string) int {
	input, err := utils.ReadInput(filename)
	solution := 0

	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	directions := strings.Split(input[0], "")
	allPaths := []Path{}
	pathMap := map[string]int{}

	for i := 2; i < len(input); i++ {
		pathways := strings.Split(input[i], "=")
		path := strings.TrimSpace(pathways[0])
		directions := strings.Split(pathways[1], ",")
		leftPath := strings.TrimSpace(strings.Replace(directions[0], "(", "", 1))
		rightPath := strings.TrimSpace(strings.Replace(directions[1], ")", "", 1))
		newPath := Path{path: path, left: leftPath, right: rightPath}
		allPaths = append(allPaths, newPath)
		pathMap[newPath.path] = i - 2
	}

	directionIdx := 0
	nextPathIdx := 0
	nextPath := allPaths[0]

	for nextPath.path != END_PATH {
		direction := directions[directionIdx%len(directions)]
		directionIdx++
		solution++

		if direction == DIRECTION_LEFT {
			nextPathIdx = pathMap[nextPath.left]
		} else {
			nextPathIdx = pathMap[nextPath.right]
		}
		nextPath = allPaths[nextPathIdx]
	}

	return solution
}

func (d Day) Part2(filename string) int {
	_, err := utils.ReadInput(filename)
	solution := 0

	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	return solution
}
