package main

import (
	"adventofcode2023/utils"
	"fmt"
	"log"
	"strings"
)

const (
	sampleInput1 = "test_input1.txt"
	sampleInput2 = "test_input2.txt"
	puzzleInput  = "input.txt"
)

type Day struct{}

type Path struct {
	path  string
	right string
	left  string
}

type GhostPath struct {
	startPath string
	nextPath  string
	left      string
	right     string
	reachedZ  bool
	steps     int
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

	solution2 := day.Part2(puzzleInput)
	fmt.Printf("Part 2 solution %d\n", solution2)
}

func (d Day) Part1(filename string) int {
	input, err := utils.ReadInput(filename)
	solution := 0

	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	directions := strings.Split(input[0], "")
	allPaths := map[string]Path{}

	for i := 2; i < len(input); i++ {
		pathways := strings.Split(input[i], "=")
		path := strings.TrimSpace(pathways[0])
		directions := strings.Split(pathways[1], ",")
		leftPath := strings.TrimSpace(strings.Replace(directions[0], "(", "", 1))
		rightPath := strings.TrimSpace(strings.Replace(directions[1], ")", "", 1))
		newPath := Path{path: path, left: leftPath, right: rightPath}
		allPaths[newPath.path] = newPath
	}

	directionIdx := 0
	currPath := "AAA"

	for currPath != END_PATH {
		direction := directions[directionIdx%len(directions)]
		directionIdx++
		solution++

		if direction == DIRECTION_LEFT {
			currPath = allPaths[currPath].left
		} else {
			currPath = allPaths[currPath].right
		}
	}

	return solution
}

func (d Day) Part2(filename string) int {
	input, err := utils.ReadInput(filename)
	solution := 0

	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	directions := strings.Split(input[0], "")
	allPaths := map[string]Path{}
	ghostPaths := []GhostPath{}

	for i := 2; i < len(input); i++ {
		pathways := strings.Split(input[i], "=")
		path := strings.TrimSpace(pathways[0])
		directions := strings.Split(pathways[1], ",")
		leftPath := strings.TrimSpace(strings.Replace(directions[0], "(", "", 1))
		rightPath := strings.TrimSpace(strings.Replace(directions[1], ")", "", 1))
		newPath := Path{path: path, left: leftPath, right: rightPath}
		allPaths[path] = newPath
		if strings.HasSuffix(path, "A") {
			ghostPaths = append(ghostPaths, GhostPath{
				startPath: path,
				nextPath:  path,
				left:      leftPath,
				right:     rightPath,
				reachedZ:  false,
				steps:     0,
			})
		}
	}

	directionIdx := 0

	for !d.ghostPathsEnded(ghostPaths) {
		direction := directions[directionIdx%len(directions)]
		directionIdx++

		for idx, gNode := range ghostPaths {
			if gNode.reachedZ {
				continue
			}

			if direction == DIRECTION_LEFT {
				ghostPaths[idx].nextPath = allPaths[gNode.nextPath].left
			} else {
				ghostPaths[idx].nextPath = allPaths[gNode.nextPath].right
			}

			if strings.HasSuffix(gNode.nextPath, "Z") {
				ghostPaths[idx].reachedZ = true
				ghostPaths[idx].steps = solution
			}
		}
		solution++
	}

	steps := []int{}
	for _, gP := range ghostPaths {
		steps = append(steps, gP.steps)
	}

	//https://www.calculator.net/lcm-calculator.html?numberinputs=12169%2C20093%2C20659%2C22357%2C13301%2C18961&x=Calculate

	return solution
}

func (d Day) ghostPathsEnded(paths []GhostPath) bool {
	for _, gP := range paths {
		if !gP.reachedZ {
			return false
		}
	}

	return true
}
